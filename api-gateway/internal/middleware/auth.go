package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			return
		}

		token := parts[1]
		if len(token) < 10 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token length"})
			return
		}

		// In a real application, you would validate the token with your auth service
		// For now, we'll just set the user_id from the token
		// This is just for demonstration, in production use proper JWT validation
		c.Set("user_id", "user-123")
		c.Next()
	}
}
