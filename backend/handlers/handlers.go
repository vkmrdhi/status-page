package handlers

import (
	"backend/models"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Create a new team
func CreateTeam(c *gin.Context) {
	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, team)
}

// Get all teams
func GetTeams(c *gin.Context) {
	var teams []models.Team
	if err := models.DB.Find(&teams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

// Get a specific team
func GetTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team
	if err := models.DB.First(&team, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, team)
}

// Update a specific team
func UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team
	if err := models.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, team)
}

// Delete a specific team
func DeleteTeam(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Team{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
}

// CRUD operations for Organization
func CreateOrganization(c *gin.Context) {
	var organization models.Organization
	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, organization)
}

func GetOrganizations(c *gin.Context) {
	var organizations []models.Organization
	if err := models.DB.Find(&organizations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, organizations)
}

func GetOrganization(c *gin.Context) {
	id := c.Param("id")
	var organization models.Organization
	if err := models.DB.First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}
	c.JSON(http.StatusOK, organization)
}

func UpdateOrganization(c *gin.Context) {
	id := c.Param("id")
	var organization models.Organization
	if err := models.DB.First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Organization not found"})
		return
	}
	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, organization)
}

func DeleteOrganization(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Organization{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Organization deleted"})
}

// CRUD operations for Services
func CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, service)
}

func GetServices(c *gin.Context) {
	var services []models.Service
	if err := models.DB.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

func GetService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	if err := models.DB.First(&service, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, service)
}

func UpdateService(c *gin.Context) {
	type StatusUpdateRequest struct {
		Status string `json:"status"`
	}

	var req StatusUpdateRequest
	serviceID := c.Param("id") // Service ID from the route

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var service models.Service
	if err := models.DB.First(&service, serviceID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	// Update service status
	previousStatus := service.Status
	service.Status = req.Status
	if err := models.DB.Save(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update service"})
		return
	}

	// Handle incidents based on status change
	if shouldCreateIncident(previousStatus, req.Status) {
		incident := models.Incident{
			Title:       "Service Issue Detected",
			Description: fmt.Sprintf("The %s has entered a degraded or outage state.", service.Name),
			Status:      "active",
			Priority:    getIncidentPriority(req.Status), // Set priority based on the status
			ServiceID:   service.ID,
			CreatedAt:   time.Now(),
		}
		if err := models.DB.Create(&incident).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create incident"})
			return
		}
	} else if shouldResolveIncident(previousStatus, req.Status) {
		var activeIncident models.Incident
		if err := models.DB.Where("service_id = ? AND status = ?", service.ID, "active").First(&activeIncident).Error; err == nil {
			activeIncident.Status = "resolved"
			activeIncident.ResolvedAt = timePtr(time.Now())
			if err := models.DB.Save(&activeIncident).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to resolve incident"})
				return
			}
		}
	}

	c.JSON(http.StatusOK, service)
}

// Helper functions to determine incident logic
func shouldCreateIncident(previousStatus, newStatus string) bool {
	return previousStatus == "operational" &&
		(newStatus == "degraded" || newStatus == "partial_outage" || newStatus == "major_outage")
}

func shouldResolveIncident(previousStatus, newStatus string) bool {
	return (previousStatus == "degraded" || previousStatus == "partial_outage" || previousStatus == "major_outage") &&
		newStatus == "operational"
}

// Helper function to determine priority based on the service status
func getIncidentPriority(status string) string {
	switch status {
	case "degraded":
		return "medium"
	case "partial_outage":
		return "high"
	case "major_outage":
		return "critical"
	default:
		return "low" // Default to low if operational or other status
	}
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func DeleteService(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Service{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted"})
}

// CRUD operations for Incidents
func CreateIncident(c *gin.Context) {
	var incident models.Incident
	if err := c.ShouldBindJSON(&incident); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&incident).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, incident)
}

func GetIncidents(c *gin.Context) {
	var incidents []models.Incident

	if err := models.DB.Find(&incidents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sort.Slice(incidents, func(i, j int) bool {
		// Compare the CreatedAt timestamps
		return incidents[i].CreatedAt.After(incidents[j].CreatedAt)
	})

	// Return the sorted incidents
	c.JSON(http.StatusOK, incidents)
}

func GetIncident(c *gin.Context) {
	id := c.Param("id")
	var incident models.Incident
	if err := models.DB.First(&incident, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incident not found"})
		return
	}
	c.JSON(http.StatusOK, incident)
}

// Update an incident
func UpdateIncident(c *gin.Context) {
	id := c.Param("id")
	var incident models.Incident
	if err := models.DB.First(&incident, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incident not found"})
		return
	}
	if err := c.ShouldBindJSON(&incident); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Save(&incident).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, incident)
}

// Delete an incident
func DeleteIncident(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Incident{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Incident deleted"})
}
