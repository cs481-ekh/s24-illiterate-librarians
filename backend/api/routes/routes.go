package routes

import (
	"LiteracyLink.com/backend/api/handlers"
	"github.com/gin-gonic/gin"
)

// Set up the routes for the application
func SetupRoutes(router *gin.Engine) {
	// Example route
	userRoutes := router.Group("/health")
	{
		userRoutes.GET("", handlers.HealthCheckHandler)
		// Add more user routes as needed
	}

	// Add more route groups or individual routes as needed
}
