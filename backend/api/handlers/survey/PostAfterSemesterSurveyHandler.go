package survey

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostAfterSemesterSurveyHandler(c *gin.Context) {
	userId := c.Param("user_id")
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to submit results for semester survey form user: %s", userId),
	})
}
