package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListServices(c *gin.Context) {
	userID := c.GetString("userID") // Extracted from token
	c.JSON(http.StatusOK, gin.H{
		"message": "Here are your services",
		"userID":  userID,
	})
}
