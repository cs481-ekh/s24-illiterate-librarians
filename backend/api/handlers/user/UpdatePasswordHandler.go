package user

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
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
	
	//get the id (therby jwt token)
	id := c.MustGet("userID");

	//set it into the passupdate struct
	request.UserID = id;

	// Validate password
	if !isValidPassword(request.PasswordHash) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	// Hash the password
	hashedPassword, err := hashPassword(request.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Set the hashed password in the User struct
	request.PasswordHash = hashedPassword

	// Now you can use the request to update the password
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

	// return 204 to front end showing success
	return c.JSON(http.StatusNoContent, gin.H{})

}


