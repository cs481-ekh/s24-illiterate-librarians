package model

import "time"

type EOSParentSurvey struct {
	EOSPSID            []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"EOS_p_s_id"`
	ChildID            []byte    `gorm:"type:BINARY(16);not null" json:"child_id"`
	ParentID           []byte    `gorm:"type:BINARY(16);not null" json:"parent_id"`
	SemesterID         []byte    `gorm:"type:BINARY(16);not null" json:"semester_id"`
	SurveyCompleteDate time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"survey_complete_date"`

	// Add other fields to represent the responses recorded in the survey
	ChildBarrierReading string    `gorm:"type:VARCHAR(511);not null" json:"child_barrier_reading"`
	ChildBarrierWriting string    `gorm:"type:VARCHAR(511);not null" json:"child_barrier_writing"`
	WantParentTraining  int8      `gorm:"type:TINYINT;not null" json:"want_parent_training"` // 1 (very likely), 2 (Likely), 3 (Unsure), 4 (unlikely), 5 (very unlikely)

	OnlineModules      bool      `gorm:"default:false" json:"online_modules"`
	ZoomMeetings       bool      `gorm:"default:false" json:"zoom_meetings"`
	InPerson           bool      `gorm:"default:false" json:"in_person"`
	Blended            bool      `gorm:"default:false" json:"blended"`
	IndividualCoaching bool      `gorm:"default:false" json:"individual_coaching"`

	FamilyTutorRelationship    int8      `gorm:"type:TINYINT;not null" json:"family_tutor_relationship"`
	FamilyTtutorCommunication  int8      `gorm:"type:TINYINT;not null" json:"family_tutor_communication"`
	ChildInstructionRecieved   int8      `gorm:"type:TINYINT;not null" json:"child_instruction_recieved"`
	ChildEnjoyment             int8      `gorm:"type:TINYINT;not null" json:"child_enjoyment"`
	ChildConfidenceR           int8      `gorm:"type:TINYINT;not null" json:"child_confidence_r"` //confidence in reading
	ChildConfidenceW         int8      `gorm:"type:TINYINT;not null" json:"child_confidence_w"` //writing
	ChildConfidenceS         int8      `gorm:"type:TINYINT;not null" json:"child_confidence_s"` //spelling

	PreferZoom                bool      `gorm:"default:false;not null" json:"prefer_zoom"` //true if prefer online(zoom), false if prefer in-person (college of education building)
	ChildEnjoyMost           string    `gorm:"type:VARCHAR(511);not null" json:"child_enjoy_most"`
	ImprovmentsRecommendation string    `gorm:"type:VARCHAR(511);not null" json:"improvments_recommendation"`
	MiscFeedback              string    `gorm:"type:VARCHAR(511);not null" json:"misc_feedback"`

	// Define fields for relationships with other tables
	Child    Child    `gorm:"foreignKey:ChildID"`
	Parent   Parent   `gorm:"foreignKey:ParentID"`
	Semester Semester `gorm:"foreignKey:SemesterID"`
}
