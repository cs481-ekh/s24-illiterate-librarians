package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LookupAllUsersHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "TODO make func return all users",
	})
}
