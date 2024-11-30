package routes

import (
	"net/http"

	"github.com/vikasatfactors/status-page-app/backend/internal/handlers"
	"github.com/vikasatfactors/status-page-app/backend/internal/middleware"

	IWS "github.com/vikasatfactors/status-page-app/backend/internal/websocket"

	"github.com/gin-gonic/gin"
	WS "github.com/gorilla/websocket"
)

// WebSocketUpgrader for handling WebSocket connections
var upgrader = WS.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// In production, implement proper origin checking
		return true
	},
}

// RouterConfig holds all the handlers and dependencies for routing
type RouterConfig struct {
	UserHandler         *handlers.UserHandler
	ServiceHandler      *handlers.ServiceHandler
	IncidentHandler     *handlers.IncidentHandler
	MaintenanceHandler  *handlers.MaintenanceHandler
	PublicStatusHandler *handlers.PublicStatusHandler
	WebSocketManager    *IWS.Manager
	AuthMiddleware      *middleware.AuthMiddleware
}

// SetupRouter configures and returns a Gin router with all routes
func SetupRouter(config *RouterConfig) *gin.Engine {
	router := gin.Default()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	// Public routes
	public := router.Group("/api/public")
	{
		// Public status page routes
		public.GET("/status/:subdomain", config.PublicStatusHandler.GetPublicStatus)

		// WebSocket route for public status updates
		public.GET("/ws/:subdomain", func(c *gin.Context) {
			subdomain := c.Param("subdomain")

			// Find organization by subdomain
			org, err := config.AuthMiddleware.GetOrganizationBySubdomain(subdomain)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
				return
			}

			// Upgrade to WebSocket
			conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
				return
			}
			defer conn.Close()

			// Handle WebSocket connection
			config.WebSocketManager.HandleWebSocketConnection(conn, org.ID)
		})
	}

	// Authentication routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", config.UserHandler.CreateUser)
		// Add login, logout routes as needed
	}

	// Protected routes (require authentication)
	protected := router.Group("/api")
	protected.Use(config.AuthMiddleware.AuthenticateUser())
	{
		// Organization-level middleware to ensure user is in the right org
		protected.Use(config.AuthMiddleware.OrganizationAccess())

		// Service routes
		services := protected.Group("/services")
		{
			services.POST("", config.ServiceHandler.CreateService)
			services.PUT("/:id/status", config.ServiceHandler.UpdateServiceStatus)
			services.GET("/summary", config.ServiceHandler.GetServiceStatusSummary)
		}

		// Incident routes
		incidents := protected.Group("/incidents")
		{
			incidents.POST("", config.IncidentHandler.CreateIncident)
			incidents.POST("/:id/updates", config.IncidentHandler.AddIncidentUpdate)
			incidents.PUT("/:id/resolve", config.IncidentHandler.ResolveIncident)
		}

		// Maintenance routes
		maintenances := protected.Group("/maintenances")
		{
			maintenances.POST("", config.MaintenanceHandler.ScheduleMaintenance)
			maintenances.PUT("/:id/start", config.MaintenanceHandler.StartMaintenance)
			maintenances.PUT("/:id/complete", config.MaintenanceHandler.CompleteMaintenance)
		}

		// User management routes
		users := protected.Group("/users")
		{
			users.GET("/profile", config.UserHandler.GetProfile)
			users.PUT("/profile", config.UserHandler.UpdateProfile)
		}
	}

	// Admin-only routes
	admin := router.Group("/api/admin")
	admin.Use(config.AuthMiddleware.AuthenticateUser())
	admin.Use(config.AuthMiddleware.AdminAccess())
	{
		admin.GET("/organizations", config.UserHandler.ListOrganizations)
		admin.POST("/organizations", config.UserHandler.CreateOrganization)
	}

	return router
}

// Middleware for CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Main application entry point to wire up and start the server
func RunServer(config *RouterConfig) {
	router := SetupRouter(config)

	// Configure server settings
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start WebSocket manager
	go config.WebSocketManager.Start()

	// Start HTTP server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
