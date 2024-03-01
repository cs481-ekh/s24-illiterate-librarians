package main

import (
	"LiteracyLink.com/backend/api/routes"
	"LiteracyLink.com/backend/middleware"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	// Initialize Gin router
	router := gin.Default()
	router.Use(middleware.DatabaseMiddleware())
	routes.SetupRoutes(router)
	// Set up the routes for the application
	router.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(".", "index.html"))
	})
	// Start the server
	router.Run(":8080")
}
