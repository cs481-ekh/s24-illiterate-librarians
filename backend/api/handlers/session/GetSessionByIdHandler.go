package session

import (
	"LiteracyLink.com/backend/api/model"
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
	//var response []model.TutoringSessionResponse
	sessions, err := db.GetSessionsByUserIdAndType(userID.(string), userType.(string), dbc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("could not find sessions for userid: %s", userID),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sessions": sessions,
	})
	//for _, tutorSession := range sessions {
	//	tmp := convertToTutorSessionResponse(tutorSession)
	//	response = append(response, tmp)
	//}
}

func convertToTutorSessionResponse(tutorSession model.TutorSession) model.TutoringSessionResponse {
	return model.TutoringSessionResponse{
		SessionId: tutorSession.TutorSessionID,
		Title:     "Tutoring",
		When:      tutorSession.MeetingDate,
		ZoomLink:  tutorSession.ZoomJoinLink,
	}
}
