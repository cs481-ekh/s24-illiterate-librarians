package application

import (
	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/auth"
	"LiteracyLink.com/backend/db"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

// No idea if this is this will work tbh, but it looks like it could haha

func GetTutorApplicationHandler(c *gin.Context) {
	var request model.AppRequest
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding tutor application request: %g", err),
		})
	}

	dbc := c.MustGet("db").(*gorm.DB)
	app, err := db.GetApp(request, dbc)
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
	// err = ((string(app.ParentId) != request.Parent) || (string(app.ChildId) != request.Child) || (string(app.DesiredSemesterId) != request.Semester))
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

	c.JSON(http.StatusOK, app)


}
