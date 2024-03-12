package application

import (
	"errors"
	"fmt"
	"net/http"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func UpdateTutorApplicationHandler(c *gin.Context) {
	
	var request model.TutoringApplication
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("Error while binding tutor application update request: %g", err),
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Guardian phone number"})
		return
	}

	if !isValidPhone(request.EmergencyConPhone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Emergency contact phone number"})
		return
	}

	dbc := c.MustGet("db").(*gorm.DB)
	err = db.UpdateApp(request, dbc)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "failed",
				"message": fmt.Sprintf("invalid application when trying to update"),
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
