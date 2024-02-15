package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LookupUserHandler(c *gin.Context) {
	user_id := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to lookup user with id: %s", user_id),
	})
}
