package handlers

import (
	"net/http"

	"github.com/vikasatfactors/status-page-app/backend/internal/services"

	"github.com/gin-gonic/gin"
)

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
