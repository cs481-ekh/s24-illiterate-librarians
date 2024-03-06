package model

import (
	"time"
)

// NOTE: not all of the things that will be asked in the application will
// be stored here / in the app_for_tutoring table. Some things like the days
// that parents will be unavailable will be stored in other tables. For a
// full list of questions in the application, look at the original applicaiton.

// TODO: list out in the comments here the questions that should be asked for the benefit of future developers

type Application struct {
	AppForTutId       []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"app_for_tut_id"`
	ChildId           []byte    `gorm:"type:BINARY(16);not null" json:"child_id"`
	ParentId          []byte    `gorm:"type:BINARY(16);not null" json:"parent_id"`
	AppComplete       bool      `gorm:"default:false" json:"app_complete"`
	AppCompleteDate   time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"app_complete_date"`
	DesiredSemesterId []byte    `gorm:"type:BINARY(16);not null" json:"desired_semester_id"`

	ChildDataConsent        bool `gorm:"default:false" json:"child_data_consent"`
	PhotoReleaseConsent     bool `gorm:"default:false" json:"photo_release_consent"`
	NeedFinancialAssistance bool `gorm:"default:false" json:"need_financial_assistance"`

	Guardian2FirstN string `gorm:"type:VARCHAR(50)" json:"guardian2_first_n"`
	Guardian2LastN  string `gorm:"type:VARCHAR(50)" json:"guardian2_last_n"`
	Guardian2Phone  string `gorm:"type:VARCHAR(20)" json:"guardian2_phone"`
	Guardian2Email  string `gorm:"type:VARCHAR(255);NOT NULL" json:"guardian2_email"`

	EmergencyConName     string `gorm:"type:VARCHAR(50)" json:"emergency_con_name"`
	EmergencyConRelation string `gorm:"type:VARCHAR(255)" json:"emergency_con_relation"`
	EmergencyConPhone    string `gorm:"type:VARCHAR(20)" json:"emergency_con_phone"`

	Previous_child_participation bool   `gorm:"default:false"`
	WhatSSemester                string `gorm:"type:VARCHAR(50)" json:"what_semester"`
	ChildCurrentSchool           string `gorm:"type:VARCHAR(50)" json:"child_current_school"`
	ListLanguagesSpoken          string `gorm:"type:VARCHAR(255)" json:"list_languages_spoken"`
	ReceivedSpecialEd            string `gorm:"type:VARCHAR(50)" json:"received_special_ed"`
	ListChallenges               string `gorm:"type:VARCHAR(50)" json:"list_challenges"`
	HowLongConcerned             string `gorm:"type:VARCHAR(50)" json:"how_long_concerned"`
	DescribeHopes                string `gorm:"type:VARCHAR(50)" json:"describe_hopes"`
	ChildAllergyMeds             string `gorm:"type:VARCHAR(50)" json:"child_allergy_meds"`
	MiscInfo                     string `gorm:"type:VARCHAR(50)" json:"misc_info"`
	HearAboutLitLab              string `gorm:"type:VARCHAR(50)" json:"hear_about_litLab"`
}
