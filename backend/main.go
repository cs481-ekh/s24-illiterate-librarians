package main

import (
	"github.com/gin-gonic/gin"

	"LiteracyLink.com/backend/api/routes"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Serve API Routes using serve routes

	// Set up the routes for the application

	routes.SetupRoutes(router)
	routes.ServeStatic(router)

	// Start the server
	router.Run(":8080")
}
