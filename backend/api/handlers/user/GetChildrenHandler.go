package user

import (
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetChildrenHandler(c *gin.Context) {
	userID := c.MustGet("UserID").(string)
	userType := c.MustGet("UserType").(string)
	dbc := c.MustGet("db").(*gorm.DB)
	if userType != "parent" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not classified as parent",
		})
		return
	}
	children, err := db.GetChildren(userID, dbc)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, children)

}
