package system_tests

import (
	"LiteracyLink.com/backend/api/routes"
	"github.com/gin-gonic/gin"
)

func InitRouterFunction() *gin.Engine {
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	return router
}
