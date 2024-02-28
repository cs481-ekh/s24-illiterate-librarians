package model

import (
	"time"
)

type Child struct {
	ChildID   []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey"`
	UserID    []byte    `gorm:"type:BINARY(16);not null"`
	ParentID  []byte    `gorm:"type:BINARY(16);not null"`
	BirthDate time.Time `gorm:"type:DATE"`
	Grade     int8      `gorm:"type:TINYINT"`
	FirstName string    `gorm:"type:VARCHAR(50);not null"`
	LastName  string    `gorm:"type:VARCHAR(50);not null"`
	CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`

	// Define a field to represent the relationship with the Parent model
	Parent Parent `gorm:"foreignKey:ParentID"`
}
