package handlers

import (
	"net/http"
	"strconv"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"
	"github.com/vikasatfactors/status-page-app/backend/internal/services"

	"github.com/gin-gonic/gin"
)

// UserHandler manages user-related HTTP requests
type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req struct {
		Email     string `json:"email"`
		Name      string `json:"name"`
		Auth0ID   string `json:"auth0_id"`
		OrgName   string `json:"org_name"`
		Subdomain string `json:"subdomain"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUserWithOrganization(
		req.Email,
		req.Name,
		req.Auth0ID,
		req.OrgName,
		req.Subdomain,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// ServiceHandler manages service-related HTTP requests
type ServiceHandler struct {
	serviceStatusService *services.ServiceStatusService
}

func NewServiceHandler(serviceStatusService *services.ServiceStatusService) *ServiceHandler {
	return &ServiceHandler{serviceStatusService: serviceStatusService}
}

func (h *ServiceHandler) UpdateServiceStatus(c *gin.Context) {
	// Get the service ID from the URL
	serviceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
		return
	}

	// Request body for status update
	var req struct {
		Status              models.ServiceStatus `json:"status"`
		IncidentTitle       string               `json:"incident_title,omitempty"`
		IncidentDescription string               `json:"incident_description,omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare incident details if not operational
	var incident *models.Incident
	if req.Status != models.StatusOperational {
		incident = &models.Incident{
			Title:       req.IncidentTitle,
			Description: req.IncidentDescription,
		}
	}

	// Update service status
	if err := h.serviceStatusService.UpdateServiceStatus(
		uint(serviceID),
		req.Status,
		incident,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service status updated successfully"})
}

func (h *ServiceHandler) GetServiceStatusSummary(c *gin.Context) {
	// Get organization ID from context (set by middleware)
	orgID, exists := c.Get("organization_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Organization not found"})
		return
	}

	summary, err := h.serviceStatusService.GetServiceStatusSummary(orgID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}

// IncidentHandler manages incident-related HTTP requests
type IncidentHandler struct {
	incidentService *services.IncidentService
}

func NewIncidentHandler(incidentService *services.IncidentService) *IncidentHandler {
	return &IncidentHandler{incidentService: incidentService}
}

func (h *IncidentHandler) CreateIncident(c *gin.Context) {
	var incident models.Incident
	if err := c.ShouldBindJSON(&incident); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get organization ID from context
	orgID, exists := c.Get("organization_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Organization not found"})
		return
	}
	incident.OrganizationID = orgID.(uint)

	if err := h.incidentService.CreateIncident(&incident); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, incident)
}

func (h *IncidentHandler) AddIncidentUpdate(c *gin.Context) {
	// Get incident ID from URL
	incidentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid incident ID"})
		return
	}

	var update models.IncidentUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.incidentService.AddIncidentUpdate(uint(incidentID), &update); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Incident update added successfully"})
}

func (h *IncidentHandler) ResolveIncident(c *gin.Context) {
	// Get incident ID from URL
	incidentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid incident ID"})
		return
	}

	if err := h.incidentService.ResolveIncident(uint(incidentID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Incident resolved successfully"})
}

// MaintenanceHandler manages maintenance-related HTTP requests
type MaintenanceHandler struct {
	maintenanceService *services.MaintenanceService
}

func NewMaintenanceHandler(maintenanceService *services.MaintenanceService) *MaintenanceHandler {
	return &MaintenanceHandler{maintenanceService: maintenanceService}
}

func (h *MaintenanceHandler) ScheduleMaintenance(c *gin.Context) {
	var maintenance models.Maintenance
	if err := c.ShouldBindJSON(&maintenance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get organization ID from context
	orgID, exists := c.Get("organization_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Organization not found"})
		return
	}
	maintenance.OrganizationID = orgID.(uint)

	if err := h.maintenanceService.ScheduleMaintenance(&maintenance); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, maintenance)
}

func (h *MaintenanceHandler) StartMaintenance(c *gin.Context) {
	// Get maintenance ID from URL
	maintenanceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maintenance ID"})
		return
	}

	if err := h.maintenanceService.StartMaintenance(uint(maintenanceID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Maintenance started successfully"})
}

func (h *MaintenanceHandler) CompleteMaintenance(c *gin.Context) {
	// Get maintenance ID from URL
	maintenanceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maintenance ID"})
		return
	}

	if err := h.maintenanceService.CompleteMaintenance(uint(maintenanceID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Maintenance completed successfully"})
}

// PublicStatusHandler for the public-facing status page
type PublicStatusHandler struct {
	serviceStatusService *services.ServiceStatusService
}

func NewPublicStatusHandler(serviceStatusService *services.ServiceStatusService) *PublicStatusHandler {
	return &PublicStatusHandler{serviceStatusService: serviceStatusService}
}

func (h *PublicStatusHandler) GetPublicStatus(c *gin.Context) {
	// Get subdomain from URL
	subdomain := c.Param("subdomain")

	// Fetch organization by subdomain
	org, err := h.organizationRepo.FindBySubdomain(subdomain)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}

	// Get service status summary
	summary, err := h.serviceStatusService.GetServiceStatusSummary(org.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, summary)
}
