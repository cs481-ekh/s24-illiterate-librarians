package event

import (
	"LiteracyLink.com/backend/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetAnnouncementsHandler(c *gin.Context) {
	dbc := c.MustGet("db").(*gorm.DB)
	response, err := db.GetAnnouncements(dbc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("Error while getting announcements: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
