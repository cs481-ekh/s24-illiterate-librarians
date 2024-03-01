package main

import (
	"LiteracyLink.com/backend/api/routes"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	// Initialize Gin router
	router := gin.Default()
	// Set up the routes for the application
	routes.SetupRoutes(router)
	router.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(".", "index.html"))
	})
	// Start the server
	router.Run(":8080")
}
