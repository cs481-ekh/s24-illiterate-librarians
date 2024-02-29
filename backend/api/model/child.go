package model

import (
	"time"
)

type Child struct {
	ChildID   []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"child_id"`
	ParentID  []byte    `gorm:"type:BINARY(16);not null" json:"parent_id"`
	BirthDate time.Time `gorm:"type:DATE" json:"birth_date"`
	Grade     int8      `gorm:"type:TINYINT" json:"grade"`
	FirstName string    `gorm:"type:VARCHAR(50);not null" json:"first_name"`
	LastName  string    `gorm:"type:VARCHAR(50);not null" json:"last_name"`
	CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`

	// Define a field to represent the relationship with the Parent model
	Parent Parent `gorm:"foreignKey:ParentID"`
}
