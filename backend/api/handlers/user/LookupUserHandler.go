package user

import (
	"errors"
	"fmt"
	"net/http"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/google/uuid"

)

func LookupUserHandler(c *gin.Context) {
	var request model.UserRequest
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding user_id with user object request: %g", err),
		})
	}

	dbc := c.MustGet("db").(*gorm.DB)
	user, err := db.GetUser(request, dbc)
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


	// Parse UserID as UUID
	UID, err := uuid.Parse(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID format"})
		return
	}

	if (user.UserID != UID) {
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

	c.JSON(http.StatusOK, user)
}
