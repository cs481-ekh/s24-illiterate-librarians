package application

import (
	"errors"
	"fmt"
	"net/http"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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

	if (string(app.ParentId) != request.Parent) || (string(app.ChildId) != request.Child) || (string(app.DesiredSemesterId) != request.Semester) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("wrong indentifying info"),
		})
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("ERROR: %g", err),
		})
		return
	}

	c.JSON(http.StatusOK, app)

}
