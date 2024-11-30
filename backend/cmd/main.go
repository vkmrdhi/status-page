package main

import (
	"log"

	"github.com/vikasatfactors/status-page-app/backend/internal/config"
	"github.com/vikasatfactors/status-page-app/backend/internal/handlers"
	"github.com/vikasatfactors/status-page-app/backend/internal/middleware"
	"github.com/vikasatfactors/status-page-app/backend/internal/repositories"
	"github.com/vikasatfactors/status-page-app/backend/internal/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := config.InitializeDatabase(cfg)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	// Run migrations
	if err := config.MigrateDatabase(db); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// Initialize WebSocket manager
	wsManager := websocket.NewManager()
	go wsManager.Start()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	orgRepo := repositories.NewOrganizationRepository(db)
	serviceRepo := repositories.NewServiceRepository(db)
	incidentRepo := repositories.NewIncidentRepository(db)
	maintenanceRepo := repositories.NewMaintenanceRepository(db)

	// Initialize auth middleware
	authMiddleware := middleware.NewAuthMiddleware(
		cfg.Auth0Domain,
		cfg.Auth0Audience,
	)

	// Setup Gin router
	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://your-frontend-domain.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Public routes
	public := r.Group("/api/public")
	{
		public.GET("/status", handlers.GetPublicStatus(serviceRepo, incidentRepo, maintenanceRepo))
		public.GET("/websocket", func(c *gin.Context) {
			handlers.HandleWebSocketConnection(c, wsManager)
		})
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(authMiddleware.AuthMiddleware())
	{
		// User routes
		protected.POST("/users", handlers.CreateUser(userRepo))
		protected.GET("/users/me", handlers.GetCurrentUser(userRepo))

		// Organization routes
		protected.POST("/organizations", handlers.CreateOrganization(orgRepo, userRepo))
		protected.GET("/organizations", handlers.ListOrganizations(orgRepo))

		// Service routes
		protected.POST("/services", handlers.CreateService(serviceRepo, orgRepo))
		protected.GET("/services", handlers.ListServices(serviceRepo))
		protected.PUT("/services/:id", handlers.UpdateService(serviceRepo))

		// Incident routes
		protected.POST("/incidents", handlers.CreateIncident(incidentRepo, serviceRepo, wsManager))
		protected.GET("/incidents", handlers.ListIncidents(incidentRepo))
		protected.PUT("/incidents/:id", handlers.UpdateIncident(incidentRepo, wsManager))

		// Maintenance routes
		protected.POST("/maintenance", handlers.CreateMaintenance(maintenanceRepo, serviceRepo, wsManager))
		protected.GET("/maintenance", handlers.ListMaintenances(maintenanceRepo))
		protected.PUT("/maintenance/:id", handlers.UpdateMaintenance(maintenanceRepo, wsManager))
	}

	// Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
