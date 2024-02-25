package model

import "time"

type SessionNotes struct {
	SessionNotesID []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey"`
	TutorSessionID []byte    `gorm:"type:BINARY(16);not null"`
	CreatedDate    time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`

	// Define fields for relationships with other tables
	TutorSession TutorSession `gorm:"foreignKey:TutorSessionID"`
}
