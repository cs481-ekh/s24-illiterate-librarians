package routes

import (
	"LiteracyLink.com/backend/api/handlers/admin"
	"LiteracyLink.com/backend/api/handlers/application"
	"LiteracyLink.com/backend/api/handlers/event"
	"LiteracyLink.com/backend/api/handlers/health"
	"LiteracyLink.com/backend/api/handlers/session"
	"LiteracyLink.com/backend/api/handlers/survey"
	"LiteracyLink.com/backend/api/handlers/user"
	"LiteracyLink.com/backend/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func ServeStatic(router *gin.Engine) {

	router.Use(func(c *gin.Context) {
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			c.Next()
			return
		}

		staticDir := "."
		filePath := filepath.Join(staticDir, c.Request.URL.Path)

		// Check if a file exists at the requested path
		if _, err := os.Stat(filePath); err == nil {
			c.File(filePath)
			return
		}

		// If the file doesn't exist, continue to the next handler
		c.Next()
	})

	// If no file is found, serve the index.html file
	router.NoRoute(func(c *gin.Context) {
		dir := "."
		file := "index.html"

		// Check if the file exists in the static dir
		if _, err := os.Stat(filepath.Join(dir, file)); err == nil {
			// If it exists, serve the file
			c.File(filepath.Join(dir, file))
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})
}

// Set up the routes for the application
func SetupRoutes(router *gin.Engine) {
	healthRoutes := router.Group("/api/health")
	{
		healthRoutes.GET("", health.HealthCheckHandler)
	}

	userRoutes := router.Group("/api/user")
	{
		// POST /user/create (no JWT required)
		userRoutes.POST("/create", user.CreateUserHandler)

		// GET /user/login (no JWT required)
		userRoutes.POST("/login", user.LoginHandler)

		// ... other user route definitions ...

		// Protected routes (require JWT)
		protectedRoutes := userRoutes.Group("")
		protectedRoutes.Use(middleware.AuthMiddleware())
		{
			// GET /user/lookupAll
			protectedRoutes.GET("/lookupAll", user.LookupAllUsersHandler)

			// GET /user/lookup
			protectedRoutes.GET("/lookup", user.LookupUserHandler)

			// PUT /user/update/:user_id
			protectedRoutes.PUT("/update/:userId", user.UpdateUserHandler)

			// PUT /user/update/password/:user_id
			protectedRoutes.PUT("/update/password/:userId", user.UpdatePasswordHandler)

			// DELETE /user/delete/:user_id
			protectedRoutes.DELETE("/delete/:userId", user.DeleteUserHandler)
		}
	}
	surveyRoutes := router.Group("/api/survey")
	surveyRoutes.Use(middleware.AuthMiddleware())
	{

		// POST /survey/after_semester_survey/:user_id
		surveyRoutes.POST("/after_semester_survey/:userId", survey.PostAfterSemesterSurveyHandler)

		// GET /survey/after_semester_survey/taken/:user_id
		surveyRoutes.GET("/after_semester_survey/taken/:userId", survey.AfterSemesterSurveyTakenHandler)

		// GET /survey/after_semester_survey/:user_id
		surveyRoutes.GET("/after_semester_survey/:userId", survey.GetAfterSemesterSurveyHandler)
	}

	applicationRoutes := router.Group("/api/application")
	applicationRoutes.Use(middleware.AuthMiddleware())
	{

		// POST /survey/after_semester_survey/:user_id
		surveyRoutes.POST("/:userId", application.PostTutorApplicationHandler)

		// POST /survey/after_semester_survey/:user_id
		surveyRoutes.PUT("/:userId", application.UpdateTutorApplicationHandler)

		// GET /survey/after_semester_survey/:user_id
		surveyRoutes.GET("/:userId", application.GetTutorApplicationHandler)
	}

	sessionRoutes := router.Group("/api/session")
	sessionRoutes.Use(middleware.AuthMiddleware())
	{
		// GET /session/client/:user_id
		sessionRoutes.GET("/client/:sessionId", session.GetClientSessionsHandler)

		// GET /session/tutor/:user_id
		sessionRoutes.GET("/tutor/:sessionId", session.GetTutorSessionsHandler)

		// GET /session
		sessionRoutes.GET("", session.GetSessionByIdHandler)
	}
	eventRoutes := router.Group("/api/event")
	eventRoutes.Use(middleware.AuthMiddleware())
	{
		// GET api/event/Announcements
		eventRoutes.GET("/announcements", event.GetAnnouncementsHandler)

		// GET api/event/
		eventRoutes.GET("", event.GetEventHandler)
	}
	adminRoutes := router.Group("/api/admin")
	adminRoutes.Use(middleware.AuthMiddleware())
	{
		adminRoutes.PUT("/update/privileges", admin.UpdateUserTypeHandler)

	}

}
