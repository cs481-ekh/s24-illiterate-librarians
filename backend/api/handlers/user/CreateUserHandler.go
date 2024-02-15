package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUserHandler(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to create new user: %s", username),
	})
}
