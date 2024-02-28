package model

import "time"

type TutorSession struct {
	TutorSessionID    []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey"`
	ChildID           []byte    `gorm:"type:BINARY(16);not null"`
	ParentID          []byte    `gorm:"type:BINARY(16);not null"`
	ZoomJoinLink      string    `gorm:"type:VARCHAR(512)"`
	ZoomRecordingLink string    `gorm:"type:VARCHAR(512)"`
	MeetingDate       time.Time `gorm:"type:DATETIME;not null"`
	ParentAvail       bool      `gorm:"default:true"`
	TutorID           []byte    `gorm:"type:BINARY(16);not null"`
	SemesterID        []byte    `gorm:"type:BINARY(16);not null"`

	// Define fields for relationships with other tables
	Child    Child    `gorm:"foreignKey:ChildID"`
	Parent   Parent   `gorm:"foreignKey:ParentID"`
	Tutor    Tutor    `gorm:"foreignKey:TutorID"`
	Semester Semester `gorm:"foreignKey:SemesterID"`
}
