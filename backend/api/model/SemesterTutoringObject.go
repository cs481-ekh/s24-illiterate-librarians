package model

import "time"

type SemesterTutoringObj struct {
	SemesterTutoringID      []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey"`
	ChildID                 []byte    `gorm:"type:BINARY(16);not null"`
	ParentID                []byte    `gorm:"type:BINARY(16);not null"`
	ApplicationApproved     bool      `gorm:"default:false"`
	EOSParentSurveyComplete bool      `gorm:"default:false"`
	SurveyCompleteDate      time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	EOSParentSurveyID       []byte    // Foreign key

	EOSReportPosted   bool   `gorm:"default:false"`
	EOSReportFilePath string `gorm:"type:VARCHAR(127)"`
	SemesterID        []byte `gorm:"type:BINARY(16);not null"`

	// Define fields for relationships with other tables
	Child           Child           `gorm:"foreignKey:ChildID"`
	Parent          Parent          `gorm:"foreignKey:ParentID"`
	EOSParentSurvey EOSParentSurvey `gorm:"foreignKey:EOSParentSurveyID"`
	Semester        Semester        `gorm:"foreignKey:SemesterID"`
}
