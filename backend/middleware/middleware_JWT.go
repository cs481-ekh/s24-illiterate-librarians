package middleware

import (
	"log"

	"LiteracyLink.com/backend/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		_, err := auth.VerifyUserToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// The user is authenticated, continue to the next middleware or handler
		entry := log.Default()
		entry.Print("Token Validated")
		c.Next()
	}
}
