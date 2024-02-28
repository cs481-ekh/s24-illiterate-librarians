package main

import (
	"LiteracyLink.com/backend/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()
	// Set up routes
	routes.SetupRoutes(router)
	// Start the server
	router.Run(":8080")
}
