package middleware

import (
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Database struct to hold the GORM database connection
type Database struct {
	Conn *gorm.DB
}

// Middleware function to set up the GORM database connection
func DatabaseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Connect to the database
		database := db.ConnectDB()

		// Set the GORM database connection into the Gin context
		c.Set("db", database)

		// Pass control to the next middleware or route handler
		c.Next()

		// Disconnect from the database after the request is processed
		db.DisconnectDB(database)
	}
}
