package model

import "time"

type EOSParentSurvey struct {
	EOSPSID            []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey"`
	ChildID            []byte    `gorm:"type:BINARY(16);not null"`
	ParentID           []byte    `gorm:"type:BINARY(16);not null"`
	SemesterID         []byte    `gorm:"type:BINARY(16);not null"`
	SurveyCompleteDate time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`

	// Add other fields to represent the responses recorded in the survey

	// Define fields for relationships with other tables
	Child    Child    `gorm:"foreignKey:ChildID"`
	Parent   Parent   `gorm:"foreignKey:ParentID"`
	Semester Semester `gorm:"foreignKey:SemesterID"`
}
