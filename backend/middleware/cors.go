// middleware/cors.go
package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSConfig sets up CORS options for the application
func CORSConfig() gin.HandlerFunc {
	// Set up CORS with the desired options
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
}
