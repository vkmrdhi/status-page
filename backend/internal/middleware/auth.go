package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"
	"github.com/vikasatfactors/status-page-app/backend/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware handles authentication and authorization
type AuthMiddleware struct {
	userRepo         *repositories.UserRepository
	organizationRepo *repositories.OrganizationRepository
}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware(userRepo *repositories.UserRepository, orgRepo *repositories.OrganizationRepository) *AuthMiddleware {
	return &AuthMiddleware{
		userRepo:         userRepo,
		organizationRepo: orgRepo,
	}
}

// AuthenticateUser middleware verifies JWT token
func (m *AuthMiddleware) AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Extract token (expecting "Bearer <token>")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Verify token (this is a placeholder - replace with actual Auth0 token verification)
		claims, err := verifyToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Fetch user by Auth0 ID
		user, err := m.userRepo.FindByAuth0ID(claims.Auth0ID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Set user and organization context for downstream handlers
		c.Set("user_id", user.ID)
		c.Set("organization_id", user.OrganizationID)
		c.Set("user_role", user.Role)

		c.Next()
	}
}

// OrganizationAccess middleware ensures user is in the correct organization
func (m *AuthMiddleware) OrganizationAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user's organization ID from context
		_, exists := c.Get("organization_id")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Organization context not found"})
			c.Abort()
			return
		}

		// Additional organization-level checks can be added here
		c.Next()
	}
}

// AdminAccess middleware ensures only admin users can access certain routes
func (m *AuthMiddleware) AdminAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user role from context
		role, exists := c.Get("user_role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "User role not found"})
			c.Abort()
			return
		}

		// Check if user is an admin
		if role.(models.UserRole) != models.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// TokenClaims represents the structure of JWT claims
type TokenClaims struct {
	jwt.StandardClaims
	Auth0ID string `json:"sub"`
	Email   string `json:"email"`
}

// verifyToken is a placeholder for token verification
func verifyToken(tokenString string) (*TokenClaims, error) {
	// In a real implementation, this would:
	// 1. Verify the token with Auth0's public key
	// 2. Check token expiration
	// 3. Validate token integrity

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Placeholder secret key - replace with actual Auth0 verification
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
