package middleware

import (
	"log"

	"LiteracyLink.com/backend/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			c.JSON(401, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}
		tokenString := cookie.Value

		userId, userType, err := auth.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// The user is authenticated, continue to the next middleware or handler
		entry := log.Default()
		entry.Printf("Token Validated userID: %s , userType: %s", userId, userType)
		c.Set("UserID", userId)
		c.Set("UserType", userType)
		c.Next()
	}
}
