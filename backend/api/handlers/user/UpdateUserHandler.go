package user

import (
	"LiteracyLink.com/backend/api/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUserHandler(c *gin.Context) {
	user_id := c.Param("user_id")
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "error binding json body to user type"})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func let user with id: %s update password", user_id),
	})
}
