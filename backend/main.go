package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"LiteracyLink.com/backend/api/routes"
	"LiteracyLink.com/backend/db"
	"LiteracyLink.com/backend/middleware"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Serve API Routes using serve routes

	// Set up the routes for the application

	connection := db.ConnectDB()
	router.Use(middleware.DatabaseMiddleware(connection))

	// CORS middleware
	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"http://localhost:8080", "http://localhost:4200", "https://sdp.boisestate.edu/s24-illiterate-librarians"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowCredentials: true,
			AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
			MaxAge:           12 * time.Hour,
		}))
	routes.SetupRoutes(router)
	routes.ServeStatic(router)

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
