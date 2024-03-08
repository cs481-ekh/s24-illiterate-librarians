package application

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func PostTutorApplicationHandler(c *gin.Context) {
	
	var request model.TutoringApplication
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding tutor application post request: %g", err),
		})
	}

	//validate max char lengths
	err = ensureMaxChar(50, request.Guardian2FirstN)
	err = ensureMaxChar(50, request.Guardian2LastN)
	err = ensureMaxChar(20, request.Guardian2Phone)
	err = ensureMaxChar(255, request.Guardian2Email)

	err = ensureMaxChar(50, request.EmergencyConName)
	err = ensureMaxChar(255, request.EmergencyConRelation)
	err = ensureMaxChar(20, request.EmergencyConPhone)

	err = ensureMaxChar(50, request.WhatSSemester)
	err = ensureMaxChar(50, request.ChildCurrentSchool)
	err = ensureMaxChar(255, request.ListLanguagesSpoken)
	err = ensureMaxChar(511, request.ReceivedSpecialEd)
	err = ensureMaxChar(511, request.ListChallenges)
	err = ensureMaxChar(255, request.HowLongConcerned)
	err = ensureMaxChar(511, request.DescribeHopes)
	err = ensureMaxChar(255, request.ChildAllergyMeds)
	err = ensureMaxChar(511, request.MiscInfo)
	err = ensureMaxChar(255, request.HearAboutLitLab)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("Invalid entry length"),
		})
		return
	}

	if !isValidEmail(request.Guardian2Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}
	
	if !isValidPhone(request.Guardian2Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	}

	if !isValidPhone(request.EmergencyConPhone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	}

	dbc := c.MustGet("db").(*gorm.DB)
	err = db.SubmitApp(request, dbc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": fmt.Sprintf("invalid application when trying to post"),
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

	c.JSON(http.StatusOK)
}


func ensureMaxChar(leng int, variable string) (error) {
	if len(variable) > leng || len(variable) < 1 {
		return errors.New("Invalid entry length")
	}
	return nil
}

func isValidEmail(email string) bool {
	// Use a regular expression to validate the email address
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidPhone(phone string) bool {
	// Use a regular expression to validate the email address
	phoneRegex := regexp.MustCompile(`(?:^|[^0-9])(1[34578][0-9]{9})(?:$|[^0-9])`)
	return phoneRegex.MatchString(phone)
}