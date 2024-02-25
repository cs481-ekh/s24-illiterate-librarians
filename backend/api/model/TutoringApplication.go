package model

import "time"

type TutoringApplication struct {
	AppForTutID       []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey"`
	ChildID           []byte    `gorm:"type:BINARY(16);not null"`
	ParentID          []byte    `gorm:"type:BINARY(16);not null"`
	AppComplete       bool      `gorm:"default:false"`
	AppCompleteDate   time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	DesiredSemesterID []byte    `gorm:"type:BINARY(16);not null"`

	// Add other fields to represent the responses recorded for the app

	// NOTE: guardian 2 is not required to be filled out
	Guardian2First string `gorm:"type:VARCHAR(50)"`
	Guardian2Last  string `gorm:"type:VARCHAR(50)"`
	Guardian2Phone string `gorm:"type:VARCHAR(20)"`
	Guardian2Email string `gorm:"type:VARCHAR(255)"`

	EmergencyConName     string `gorm:"type:VARCHAR(50);not null"`
	EmergencyConRelation string `gorm:"type:VARCHAR(255);not null"`
	EmergencyConPhone    string `gorm:"type:VARCHAR(20);not null"`

	PreviousChildParticipation bool   `gorm:"default:false"`
	WhatSemester               string `gorm:"type:VARCHAR(50)"`
	ChildCurrentSchool         string `gorm:"type:VARCHAR(50);not null"`
	ListLanguagesSpoken        string `gorm:"type:VARCHAR(255);not null"`
	ReceivedSpecialEd          string `gorm:"type:VARCHAR(511);not null"`
	ListChallenges             string `gorm:"type:VARCHAR(511);not null"`
	HowLongConcerned           string `gorm:"type:VARCHAR(255);not null"`
	DescribeHopes              string `gorm:"type:VARCHAR(511);not null"`
	ChildAllergyMeds           string `gorm:"type:VARCHAR(255);not null"`
	MiscInfo                   string `gorm:"type:VARCHAR(511)"`
	HearAboutLitLab            string `gorm:"type:VARCHAR(255)"`

	// Define fields for relationships with other tables
	Child           Child    `gorm:"foreignKey:ChildID"`
	Parent          Parent   `gorm:"foreignKey:ParentID"`
	DesiredSemester Semester `gorm:"foreignKey:DesiredSemesterID"`
}
