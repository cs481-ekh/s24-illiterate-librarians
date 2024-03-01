package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTutorApplicationHandler(c *gin.Context) {
	userId := c.Param("user_id")
	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to get results for tutoring application form user: %s for this semester", userId),
	})
}