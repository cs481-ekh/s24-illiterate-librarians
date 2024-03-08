package application

import (
	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)


func PostTutorApplicationHandler(c *gin.Context) {

	//Does this attach the application that is passed in with c (ginContext) to the request model 
	//so that it is submitted to the db correctly?
	
	var request model.TutoringApplication
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding tutor application post request: %g", err),
		})
	}

	dbc := c.MustGet("db").(*gorm.DB)
	err = db.SubmitApp(request, dbc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": fmt.Sprintf("invalid application when trying to post"),
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

	c.JSON(http.StatusOK)
}
