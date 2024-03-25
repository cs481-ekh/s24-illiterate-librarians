package user

import (
	"errors"
	"fmt"
	"net/http"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUserHandler(c *gin.Context) {
	user_id := c.Param("userId")
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "error binding json body to user type"})
	}

	//is failing because of this line, how do I make the user_id variable a uuid.UUID type? I don't think type assertion will work...
	user.UserID = user_id.(uuid.UUID);


	dbc := c.MustGet("db").(*gorm.DB)
	err = db.UpdateUser(user, dbc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": fmt.Sprintf("invalid user when trying to update"),
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

	c.JSON(http.StatusOK, gin.H{})
}
