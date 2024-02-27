package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetClientSessionsHandler(c *gin.Context) {
	sessionId := c.Param("session_id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to get sessions with ID: %s", sessionId),
	})
}
