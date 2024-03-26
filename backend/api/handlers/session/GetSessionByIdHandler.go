package session

import (
	"LiteracyLink.com/backend/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetSessionByIdHandler(c *gin.Context) {
	userID, exist := c.Get("UserID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error Cound not Find UserId in request"),
		})
		return
	}
	userType, exist := c.Get("UserType")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error Cound not Find UserType in request"),
		})
		return
	}
	dbc := c.MustGet("db").(*gorm.DB)
	if userType == "parent" {
		sessions, err := db.GetClientsSessions(userID.(string), dbc)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Error: %s", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"sessions": sessions,
		})
		return
	} else if userType == "tutor" {
		sessions, err := db.GetTutorsSessions(userID.(string), dbc)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Error: %s", err),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"sessions": sessions,
		})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("could not find sessions for userid: %s", userID),
		})
		return
	}
}
