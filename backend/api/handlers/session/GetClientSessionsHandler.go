package session

import (
	"errors"
	"fmt"
	"net/http"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetClientSessionsHandler(c *gin.Context) {
	sessionId := c.Param("session_id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to get sessions with ID: %s", sessionId),
	})

	var request model.ClientSessionRequest
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding tutor session request: %g", err),
		})
	}

	dbc := c.MustGet("db").(*gorm.DB)
	ses, err := db.GetClientSession(request, dbc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": fmt.Sprintf("no application exists"),
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": fmt.Sprintf("ERROR: %g", err),
			})
			return
		}
	}

	//This error assignment is complaining because its trying to assing a bool to a type Error, TODO: fix this assignment
	// err = ((string(ses.TutorSessionID) != request.TutorSessionID) )
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  "failed",
	// 		"message": fmt.Sprintf("wrong indentifying info"),
	// 	})
	// }

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("ERROR: %g", err),
		})
		return
	}

	c.JSON(http.StatusOK, ses)
}
