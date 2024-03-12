package user

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/auth"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UpdatePasswordHandler(c *gin.Context) {
	
	var request model.PassUpdate
	err := c.BindJSON(&request)
	if err != nil {
		log.Default()
		log.Println(c.Request.Body)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding password update request: %s", err),
		})
		return
	}

	// how are we verifying the password is long enough? In the front end? I feel like it should 
	// be done in both but at this point the password is already hashed right?

	dbc := c.MustGet("db").(*gorm.DB)
	err = db.UpdatePass(request, dbc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": fmt.Sprintf("incorrect username"),
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

	//uneeded in this case?

	// err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  "failed",
	// 		"message": fmt.Sprintf("incorrect username or password"),
	// 	})
	// }

	// jwt, err := auth.GenerateJWT(user.UserID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  "failed",
	// 		"message": fmt.Sprintf("ERROR: %g", err),
	// 	})
	// 	return
	// }

	// c.SetCookie("token", jwt, 60*60*24*7, "/", "localhost", false, true)
	// c.JSON(http.StatusOK, gin.H{
	// 	"id_token":   jwt,
	// 	"expires_at": 60 * 60 * 24 * 7,
	// })
}

