package testInit

import (
	"LiteracyLink.com/backend/api/routes"
	"LiteracyLink.com/backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouterFunction() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.AuthMiddleware())
	// Set up routes
	routes.SetupRoutes(router)
	router.Use(middleware.AuthMiddleware())

	return router
}
