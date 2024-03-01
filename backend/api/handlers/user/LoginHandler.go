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
	"net/http"
)

func LoginHandler(c *gin.Context) {
	var request model.LoginRequest
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding login request: %g", err),
		})
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
		"status":  "success",
		"token":   jwt,
		"message": fmt.Sprintf("TODO: make func to let user: %s login", request.Username),
	})
}
