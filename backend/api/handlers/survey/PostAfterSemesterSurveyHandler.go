package survey

import (
	"errors"
	"fmt"
	"net/http"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func PostAfterSemesterSurveyHandler(c *gin.Context) {
	
	var request model.EOSParentSurvey
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding tutor survey post request: %g", err),
		})
	}

	//validate max char lengths
	err = ensureMaxChar(511, request.ChildBarrierReading)
	err = ensureMaxChar(511, request.ChildBarrierWriting)
	err = ensureBetween(1, 5, request.WantParentTraining)

	err = ensureBetween(1, 10, request.FamilyTutorRelationship)
	err = ensureBetween(1, 10, request.FamilyTutorCommunication)
	err = ensureBetween(1, 10, request.ChildInstructionRecieved)
	err = ensureBetween(1, 10, request.ChildEnjoyment)
	err = ensureBetween(1, 10, request.ChildConfidenceR)
	err = ensureBetween(1, 10, request.ChildConfidenceW)
	err = ensureBetween(1, 10, request.ChildConfidenceS)

	err = ensureMaxChar(511, request.ChildEnjoyMost)
	err = ensureMaxChar(511, request.ImprovmentsRecommendation)
	err = ensureMaxChar(511, request.MiscFeedback)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("Invalid entry"),
		})
		return
	}


	dbc := c.MustGet("db").(*gorm.DB)
	err = db.SubmitEOSSurvey(request, dbc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": fmt.Sprintf("invalid survey when trying to post"),
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

	c.JSON(http.StatusOK, gin.H{})
}


func ensureBetween(num1 int8, num2 int8, variable int8) (error) {
	if variable < num1 || variable > num2 {
		return errors.New("Invalid entry")
	}
	return nil
}

func ensureMaxChar(leng int, variable string) (error) {
	if len(variable) > leng || len(variable) < 1 {
		return errors.New("Invalid entry length")
	}
	return nil
}
