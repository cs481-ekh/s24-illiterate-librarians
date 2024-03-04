package user

import (
	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/auth"
	"LiteracyLink.com/backend/db"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var request model.LoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Default()
		log.Println(c.Request.Body)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding login request: %s", err),
		})
		return
	}

	dbc := c.MustGet("db").(*gorm.DB)
	user, err := db.Login(request, dbc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": fmt.Sprintf("incorrect username or password"),
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

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("incorrect username or password"),
		})
	}

	jwt, err := auth.GenerateJWT(user.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("ERROR: %g", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})
}
