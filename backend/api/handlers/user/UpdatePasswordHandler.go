package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdatePasswordHandler(c *gin.Context) {
	user_id := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func let user with id: %s update password", user_id),
	})
}
