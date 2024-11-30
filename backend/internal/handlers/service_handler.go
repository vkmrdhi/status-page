package handlers

import (
	"net/http"
	"strconv"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"
	"github.com/vikasatfactors/status-page-app/backend/internal/repositories"
	"github.com/vikasatfactors/status-page-app/backend/internal/services"

	"github.com/gin-gonic/gin"
)

// CreateService handles the creation of a new service
func CreateService(serviceRepo *repositories.ServiceRepository, orgRepo *repositories.OrganizationRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user ID from context
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		// Parse request body
		var serviceRequest struct {
			Name           string               `json:"name"`
			Description    string               `json:"description"`
			OrganizationID uint                 `json:"organization_id"`
			Status         models.ServiceStatus `json:"status"`
		}

		if err := c.BindJSON(&serviceRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate organization access
		_, err := orgRepo.GetUserOrganization(userID.(string), serviceRequest.OrganizationID)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "No access to this organization"})
			return
		}

		// Create service
		service := &models.Service{
			Name:           serviceRequest.Name,
			Description:    serviceRequest.Description,
			OrganizationID: serviceRequest.OrganizationID,
			Status:         serviceRequest.Status,
		}

		if err := serviceRepo.Create(service); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, service)
	}
}

// UpdateService handles updating an existing service
func UpdateService(serviceRepo *repositories.ServiceRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse service ID from URL
		serviceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID"})
			return
		}

		// Parse request body
		var updateRequest struct {
			Name        *string               `json:"name"`
			Description *string               `json:"description"`
			Status      *models.ServiceStatus `json:"status"`
		}

		if err := c.BindJSON(&updateRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Fetch existing service
		service, err := serviceRepo.GetByID(uint(serviceID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
			return
		}

		// Update fields if provided
		if updateRequest.Name != nil {
			service.Name = *updateRequest.Name
		}
		if updateRequest.Description != nil {
			service.Description = *updateRequest.Description
		}
		if updateRequest.Status != nil {
			service.Status = *updateRequest.Status
		}

		// Save updates
		if err := serviceRepo.Update(service); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, service)
	}
}

// ListServices retrieves services for an organization
func ListServices(serviceRepo *repositories.ServiceRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract organization ID from query parameter
		orgIDStr := c.Query("organization_id")
		orgID, err := strconv.ParseUint(orgIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization ID"})
			return
		}

		// Retrieve services
		services, err := serviceRepo.ListByOrganization(uint(orgID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, services)
	}
}

// GetPublicStatus retrieves public status of services
func GetPublicStatus(
	serviceRepo *repositories.ServiceRepository,
	incidentRepo *repositories.IncidentRepository,
	maintenanceRepo *repositories.MaintenanceRepository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve services
		services, err := serviceRepo.GetPublicStatus()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Retrieve active incidents
		incidents, err := incidentRepo.GetActiveIncidents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Retrieve active maintenances
		maintenances, err := maintenanceRepo.GetActiveMaintenances()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"services":     services,
			"incidents":    incidents,
			"maintenances": maintenances,
		})
	}
}

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
