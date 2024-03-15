package model

import "time"

type TutorSession struct {
	TutorSessionID    []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"tutor_session_id"`
	ZoomJoinLink      string    `gorm:"type:VARCHAR(512)" json:"zoom_join_link"`
	ZoomRecordingLink string    `gorm:"type:VARCHAR(512)" json:"zoom_recording_link"`
	MeetingDate       time.Time `gorm:"type:DATETIME;NOT NULL" json:"meeting_date"`
	ParentAvail       bool      `gorm:"default:true" json:"parent_avail"`
	TutorID           []byte    `gorm:"type:BINARY(16);NOT NULL" json:"tutor_id"`
	SemesterID        []byte    `gorm:"type:BINARY(16);NOT NULL" json:"semester_id"`

	// Define fields for relationships with other tables
	Tutor    Tutor    `gorm:"foreignKey:TutorID"`
	Semester Semester `gorm:"foreignKey:SemesterID"`
}

func (TutorSession) TableName() string {
	return "Tutor_session" // Specify the exact table name in the database
}
