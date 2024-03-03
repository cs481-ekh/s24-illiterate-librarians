package main

import (
	"LiteracyLink.com/backend/api/routes"
	"LiteracyLink.com/backend/db"
	"LiteracyLink.com/backend/middleware"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	// Initialize Gin router
	router := gin.Default()
	connection := db.ConnectDB()
	router.Use(middleware.DatabaseMiddleware(connection))
	routes.SetupRoutes(router)
	// Set up the routes for the application
	router.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(".", "index.html"))
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		// Close the database connection during graceful shutdown
		db.DisconnectDB(connection)
		os.Exit(0)
	}()

	// Start the server
	router.Run(":8080")
}
