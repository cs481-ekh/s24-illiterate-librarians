package user

import (
	"fmt"
	"net/http"
	"regexp"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUserHandler(c *gin.Context) {
	var userInput model.User

	// Bind the request body to the User struct
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate email
	if !isValidEmail(userInput.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	// Validate password
	if !isValidPassword(userInput.PasswordHash) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	// Hash the password
	hashedPassword, err := hashPassword(userInput.PasswordHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Set the hashed password in the User struct
	userInput.PasswordHash = hashedPassword

	// Now you can use the userInput to create a new user in the database
	// ...
	dbc := c.MustGet("db").(*gorm.DB)
	err = db.CreateUser(userInput, dbc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("ERROR: %g", err),
		})
		return
	}
	// Return success response
	c.JSON(http.StatusCreated, gin.H{})
}

func hashPassword(password string) (string, error) {
	// Use bcrypt to hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func isValidPassword(password string) bool {
	// Check password length
	if len(password) > 72 || len(password) < 8 {
		return false
	}

	// Check for at least one uppercase letter and one number
	hasUppercase := false
	hasNumber := false

	for _, char := range password {
		if 'A' <= char && char <= 'Z' {
			hasUppercase = true
		} else if '0' <= char && char <= '9' {
			hasNumber = true
		}
	}

	return hasUppercase && hasNumber
}

func isValidEmail(email string) bool {
	// Use a regular expression to validate the email address
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
