package user

import (
	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateChildHandler(c *gin.Context) {
	userID := c.MustGet("UserID").(string)
	userType := c.MustGet("UserType").(string)
	if userType != "parent" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not registered as parent",
		})
		return
	}
	dbc := c.MustGet("db").(*gorm.DB)
	var child model.Child
	err := c.BindJSON(&child)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = db.AddChild(userID, child, dbc)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
