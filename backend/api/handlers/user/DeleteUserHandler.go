package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUserHandler(c *gin.Context) {
	userId := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to delete user: %s", userId),
	})
}
