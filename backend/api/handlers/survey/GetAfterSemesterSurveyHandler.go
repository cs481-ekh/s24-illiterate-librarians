package survey

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/auth"
	"LiteracyLink.com/backend/db"
	"gorm.io/gorm"
)

func GetAfterSemesterSurveyHandler(c *gin.Context) {
	EOS_ID := c.Param("EOS_p_s_id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to get survey with ID: %s", EOS_ID),
	})

	var request model.EOSRequest
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding survey results: %g", err),
		})
	}

	
	dbc := c.MustGet("db").(*gorm.DB)
	EOS, err := db.GetEOSSurvey(request, dbc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {  
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": fmt.Sprintf("no survey exists"),
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

	err = ((string(EOS.EOSPSID) != request.EOSPSID) )
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("wrong indentifying info"),
		})
	}

	jwt, err := auth.GenerateJWT(EOS.EOSPSID)
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
		"message": fmt.Sprintf("TODO: make func to get session: %s", request.EOSPSID),
	})
}
