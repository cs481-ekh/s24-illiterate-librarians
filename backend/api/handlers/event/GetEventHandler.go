package event

import (
	"LiteracyLink.com/backend/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetEventHandler(c *gin.Context) {
	dbc := c.MustGet("db").(*gorm.DB)
	events, err := db.GetEvents(dbc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, events)
}
