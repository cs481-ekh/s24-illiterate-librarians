package main

import (
	"LiteracyLink.com/backend/api/routes"
	"LiteracyLink.com/backend/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	router.Use(middleware.AuthMiddleware())
	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
