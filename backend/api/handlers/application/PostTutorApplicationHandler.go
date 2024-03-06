package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Is a post handler even doing anything if the new application is being sent from the dao function?

func PostTutorApplicationHandler(c *gin.Context) {
	userId := c.Param("user_id")
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to submit results for tutoring application form user: %s", userId),
	})
}
