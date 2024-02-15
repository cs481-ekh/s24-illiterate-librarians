package routes

import (
	"LiteracyLink.com/backend/api/handlers/health"
	"LiteracyLink.com/backend/api/handlers/session"
	"LiteracyLink.com/backend/api/handlers/survey"
	"LiteracyLink.com/backend/api/handlers/user"
	"github.com/gin-gonic/gin"
)

// Set up the routes for the application
func SetupRoutes(router *gin.Engine) {

	healthRoutes := router.Group("/health")
	{
		healthRoutes.GET("", health.HealthCheckHandler)
	}

	userRoutes := router.Group("/user")
	{
		// POST /user/create
		userRoutes.POST("/create", user.CreateUserHandler)

		// GET /user/login
		userRoutes.GET("/login", user.LoginHandler)

		// GET /user/lookup
		userRoutes.GET("/lookup", user.LookupAllUsersHandler)

		// GET /user/lookup/:user_id
		userRoutes.GET("/lookup/:user_id", user.LookupUserHandler)

		// PUT /user/update/:user_id
		userRoutes.PUT("/update/:user_id", user.UpdateUserHandler)

		// PUT /user/update/password/:user_id
		userRoutes.PUT("/update/password/:user_id", user.UpdatePasswordHandler)

		// DELETE /user/delete/:user_id
		userRoutes.DELETE("/delete/:user_id", user.DeleteUserHandler)
	}

	surveyRoutes := router.Group("/survey")
	{
		// POST /survey/pre_semester_survey/:user_id
		surveyRoutes.POST("/pre_semester_survey/:user_id", survey.PostPreSemesterSurveyHandler)

		// POST /survey/after_semester_survey/:user_id
		surveyRoutes.POST("/after_semester_survey/:user_id", survey.PostAfterSemesterSurveyHandler)

		// GET /survey/pre_semester_survey/taken/:user_id
		surveyRoutes.GET("/pre_semester_survey/taken/:user_id", survey.PreSemesterSurveyTakenHandler)

		// GET /survey/after_semester_survey/taken/:user_id
		surveyRoutes.GET("/after_semester_survey/taken/:user_id", survey.AfterSemesterSurveyTakenHandler)

		// GET /survey/pre_semester_survey/:user_id
		surveyRoutes.GET("/pre_semester_survey/:user_id", survey.GetPreSemesterSurveyHandler)

		// GET /survey/after_semester_survey/:user_id
		surveyRoutes.GET("/after_semester_survey/:user_id", survey.GetAfterSemesterSurveyHandler)
	}

	sessionRoutes := router.Group("/session")
	{
		// GET /session/client/:user_id
		sessionRoutes.GET("/client/:user_id", session.GetClientSessionsHandler)

		// GET /session/tutor/:user_id
		sessionRoutes.GET("/tutor/:user_id", session.GetTutorSessionsHandler)

		// GET /session/:session_id
		sessionRoutes.GET("/:session_id", session.GetSessionByIdHandler)
	}

	// Add more route groups or individual routes as needed
}
