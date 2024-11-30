package handlers

import (
	"net/http"
	"strconv"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"
	"github.com/vikasatfactors/status-page-app/backend/internal/services"

	"github.com/gin-gonic/gin"
)

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
