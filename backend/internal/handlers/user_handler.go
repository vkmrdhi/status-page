package handlers

import (
	"net/http"

	"github.com/vikasatfactors/status-page-app/backend/internal/services"

	"github.com/gin-gonic/gin"
)

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
