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
	app_for_tut_id      []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey"`
	child_id            []byte    `gorm:"type:BINARY(16);not null"`
	parent_id           []byte    `gorm:"type:BINARY(16);not null"`
	app_complete        bool      `gorm:"default:false"`
	app_complete_date   time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	desired_semester_id []byte    `gorm:"type:BINARY(16);not null"`
	
	child_data_consent        bool      `gorm:"default:false"`
	photo_release_consent     bool      `gorm:"default:false"`
	need_financial_assistnace bool      `gorm:"default:false"`

	guardian2_first_n string    `gorm:"type:VARCHAR(50)" json:"guardian2_first_n"`
	guardian2_last_n  string    `gorm:"type:VARCHAR(50)" json:"guardian2_last_n"`
	guardian2_phone   string    `gorm:"type:VARCHAR(20)" json:"guardian2_phone"`
	guardian2_email   string    `gorm:"type:VARCHAR(255);NOT NULL" json:"guardian2_email"`

	emergency_con_name     string    `gorm:"type:VARCHAR(50)" json:"emergency_con_name"`
	emergency_con_relation string    `gorm:"type:VARCHAR(255)" json:"emergency_con_relation"`
	emergency_con_phone    string    `gorm:"type:VARCHAR(20)" json:"emergency_con_phone"`

	previous_child_participation bool      `gorm:"default:false"`
	what_semester                string    `gorm:"type:VARCHAR(50)" json:"what_semester"`
	child_current_school         string    `gorm:"type:VARCHAR(50)" json:"child_current_school"`
	list_languages_spoken        string    `gorm:"type:VARCHAR(255)" json:"list_languages_spoken"`
	received_special_ed          string    `gorm:"type:VARCHAR(50)" json:"received_special_ed"`
	list_challenges              string    `gorm:"type:VARCHAR(50)" json:"list_challenges"`
	how_long_concerned           string    `gorm:"type:VARCHAR(50)" json:"how_long_concerned"`
	describe_hopes               string    `gorm:"type:VARCHAR(50)" json:"describe_hopes"`
	child_allergy_meds           string    `gorm:"type:VARCHAR(50)" json:"child_allergy_meds"`
	misc_info                    string    `gorm:"type:VARCHAR(50)" json:"misc_info"`
	hear_about_litLab            string    `gorm:"type:VARCHAR(50)" json:"hear_about_litLab"`

	
}
