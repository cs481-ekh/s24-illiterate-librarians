package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserID         uuid.UUID `gorm:"type:BINARY(16);default:(UUID_TO_BIN(UUID(), 1))" json:"user_id" sql:"PRIMARY KEY"`
	Username       string    `gorm:"type:VARCHAR(255);NOT NULL;UNIQUE" json:"username"`
	PasswordHash   string    `gorm:"type:VARCHAR(255);NOT NULL" json:"password_hash"`
	Email          string    `gorm:"type:VARCHAR(255);NOT NULL;UNIQUE" json:"email"`
	PhoneNumber    string    `gorm:"type:VARCHAR(20);NOT NULL;UNIQUE" json:"phone_number"`
	FirstName      string    `gorm:"type:VARCHAR(50);NOT NULL" json:"first_name"`
	LastName       string    `gorm:"type:VARCHAR(50);NOT NULL" json:"last_name"`
	MailingAddress string    `gorm:"type:VARCHAR(255)" json:"mailing_address"`
	PrefMethodComm string    `gorm:"type:VARCHAR(255)" json:"pref_method_comm"` //Has to be "C" for call, "T" for text, or "E" for email.
	CreatedAt      time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "Users"
}
