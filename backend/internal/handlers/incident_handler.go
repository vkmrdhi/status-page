package handlers

import (
	"net/http"
	"strconv"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"
	"github.com/vikasatfactors/status-page-app/backend/internal/services"

	"github.com/gin-gonic/gin"
)

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
