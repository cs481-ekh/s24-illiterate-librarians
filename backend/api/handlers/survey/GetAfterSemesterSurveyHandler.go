package survey

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAfterSemesterSurveyHandler(c *gin.Context) {
	userId := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to get results for semester survey form for user: %s for the current semester", userId),
	})
}
