package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Database struct to hold the GORM database connection

// Middleware function to set up the GORM database connection
func DatabaseMiddleware(conn *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the GORM database connection into the Gin context
		c.Set("db", conn)

		// Pass control to the next middleware or route handler
		c.Next()

	}
}
