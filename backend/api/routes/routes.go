package routes

import (
	"LiteracyLink.com/backend/api/handlers/health"
	"LiteracyLink.com/backend/api/handlers/session"
	"LiteracyLink.com/backend/api/handlers/static"
	"LiteracyLink.com/backend/api/handlers/survey"
	"LiteracyLink.com/backend/api/handlers/user"
	"LiteracyLink.com/backend/middleware"
	"github.com/gin-gonic/gin"
)

// Set up the routes for the application
func SetupRoutes(router *gin.Engine) {

	healthRoutes := router.Group("/health")
	{
		healthRoutes.GET("", health.HealthCheckHandler)
	}

	staticRoutes := router.Group("/static")
	{
		// GET /static
		staticRoutes.GET("", static.StaticHandler)
	}

	userRoutes := router.Group("/user")
	{
		// POST /user/create (no JWT required)
		userRoutes.POST("/create", user.CreateUserHandler)

		// GET /user/login (no JWT required)
		userRoutes.GET("/login", user.LoginHandler)

		// ... other user route definitions ...

		// Protected routes (require JWT)
		protectedRoutes := userRoutes.Group("")
		protectedRoutes.Use(middleware.AuthMiddleware())
		{
			// GET /user/lookup
			protectedRoutes.GET("/lookup", user.LookupAllUsersHandler)

			// GET /user/lookup/:user_id
			protectedRoutes.GET("/lookup/:userId", user.LookupUserHandler)

			// PUT /user/update/:user_id
			protectedRoutes.PUT("/update/:userId", user.UpdateUserHandler)

			// PUT /user/update/password/:user_id
			protectedRoutes.PUT("/update/password/:userId", user.UpdatePasswordHandler)

			// DELETE /user/delete/:user_id
			protectedRoutes.DELETE("/delete/:userId", user.DeleteUserHandler)
		}
	}
	surveyRoutes := router.Group("/survey")
	surveyRoutes.Use(middleware.AuthMiddleware())
	{
		// POST /survey/pre_semester_survey/:user_id
		surveyRoutes.POST("/pre_semester_survey/:userId", survey.PostPreSemesterSurveyHandler)

		// POST /survey/after_semester_survey/:user_id
		surveyRoutes.POST("/after_semester_survey/:userId", survey.PostAfterSemesterSurveyHandler)

		// GET /survey/pre_semester_survey/taken/:user_id
		surveyRoutes.GET("/pre_semester_survey/taken/:userId", survey.PreSemesterSurveyTakenHandler)

		// GET /survey/after_semester_survey/taken/:user_id
		surveyRoutes.GET("/after_semester_survey/taken/:userId", survey.AfterSemesterSurveyTakenHandler)

		// GET /survey/pre_semester_survey/:user_id
		surveyRoutes.GET("/pre_semester_survey/:userId", survey.GetPreSemesterSurveyHandler)

		// GET /survey/after_semester_survey/:user_id
		surveyRoutes.GET("/after_semester_survey/:userId", survey.GetAfterSemesterSurveyHandler)
	}

	sessionRoutes := router.Group("/session")
	sessionRoutes.Use(middleware.AuthMiddleware())
	{
		// GET /session/client/:user_id
		sessionRoutes.GET("/client/:userId", session.GetClientSessionsHandler)

		// GET /session/tutor/:user_id
		sessionRoutes.GET("/tutor/:userId", session.GetTutorSessionsHandler)

		// GET /session/:session_id
		sessionRoutes.GET("/:sessionId", session.GetSessionByIdHandler)
	}

	// Add more route groups or individual routes as needed
}
