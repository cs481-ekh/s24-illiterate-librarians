package model

import (
	"time"
)

type Child struct {
	ChildID   []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"child_id"`
	ParentID  []byte    `gorm:"type:BINARY(16);NOT NULL" json:"parent_id"`
	BirthDate time.Time `gorm:"type:DATE;NOT NULL" json:"birth_date"`
	Grade     int8      `gorm:"type:TINYINT;NOT NULL" json:"grade"`
	FirstName string    `gorm:"type:VARCHAR(50);NOT NULL" json:"first_name"`
	LastName  string    `gorm:"type:VARCHAR(50);NOT NULL" json:"last_name"`
	CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`

	// Define a field to represent the relationship with the Parent model
	Parent Parent `gorm:"foreignKey:ParentID"`
}

type ChildJSON struct {
	ChildID   string `json:"child_id"`
	ParentID  string `json:"parent_id"`
	Grade     int8   `json:"grade"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (ChildJSON) TableName() string {
	return "Child" // Specify the exact table name in the database
}

func (Child) TableName() string {
	return "Child" // Specify the exact table name in the database
}
