package admin

import (
	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func UpdateUserTypeHandler(c *gin.Context) {
	adminID := c.MustGet("UserID").(string)
	userType := c.MustGet("UserType").(string)
	if userType != "instructor" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user token provided does not have necessary permissions for this action"})
		return
	}
	var updateUser model.UpdateUserType
	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate fields
	if updateUser.UserID == "" || updateUser.UserType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID and UserType are required"})
		return
	}

	// Parse UserID as UUID
	_, err = uuid.Parse(updateUser.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID format"})
		return
	}

	updateUser.AdminID = adminID
	dbc := c.MustGet("db").(*gorm.DB)
	err = db.SetUserType(updateUser, dbc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
