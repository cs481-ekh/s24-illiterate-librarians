package model

import "time"

type SemesterTutoringObj struct {
	SemesterTutoringID      []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"semester_tutoring_id"`
	ChildID                 []byte    `gorm:"type:BINARY(16);NOT NULL" json:"child_id"`
	ParentID                []byte    `gorm:"type:BINARY(16);NOT NULL" json:"parent_id"`
	ApplicationApproved     bool      `gorm:"default:false" json:"application_approved"`
	EOSParentSurveyComplete bool      `gorm:"default:false" json:"EOS_parent_survey_complete"`
	SurveyCompleteDate      time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"survey_complete_date"`
	EOSParentSurveyID       []byte    `gorm:"type:BINARY(16);NOT NULL" json:"EOS_parent_survey_id"`

	EOSReportPosted   bool   `gorm:"default:false" json:"EOS_report_posted"`
	EOSReportFilePath string `gorm:"type:VARCHAR(127)" json:"EOS_report_file_path"`
	SemesterID        []byte `gorm:"type:BINARY(16);NOT NULL" json:"semester_id"`

	// Define fields for relationships with other tables
	Child           Child           `gorm:"foreignKey:ChildID"`
	Parent          Parent          `gorm:"foreignKey:ParentID"`
	EOSParentSurvey EOSParentSurvey `gorm:"foreignKey:EOSParentSurveyID"`
	Semester        Semester        `gorm:"foreignKey:SemesterID"`
}
