package model

import "time"

type SessionNotes struct {
	SessionNotesID []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"session_notes_id"`
	TutorSessionID []byte    `gorm:"type:BINARY(16);NOT NULL" json:"tutor_session_id"`
	CreatedDate    time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_date"`

	// Define fields for relationships with other tables
	TutorSession TutorSession `gorm:"foreignKey:TutorSessionID"`
}
