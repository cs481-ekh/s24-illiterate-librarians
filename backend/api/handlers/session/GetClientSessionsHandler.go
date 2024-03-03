package session

import (
	"errors"
	"fmt"
	"net/http"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/auth"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {  //not sure why the gorm keyword is complaining...
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

	err = ((string(ses.TutorSessionID) != request.TutorSessionID) )
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("wrong indentifying info"),
		})
	}

	jwt, err := auth.GenerateJWT(ses.TutorSessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("ERROR: %g", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"token":   jwt,
		"message": fmt.Sprintf("TODO: make func to get session: %s", request.TutorSessionID),
	})
}
