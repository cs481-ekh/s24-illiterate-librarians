package model

import (
	"time"
)

type Semester struct {
	SemesterID      []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey"`
	SpringOrFall    string    `gorm:"type:VARCHAR(20);not null"`
	FirstSchoolDay  time.Time `gorm:"type:DATE;not null"`
	LastSchoolDay   time.Time `gorm:"type:DATE;not null"`
	FirstTutorDay   time.Time `gorm:"type:DATE;not null"`
	LastTutorDay    time.Time `gorm:"type:DATE;not null"`
	TuesdayOfBreak  time.Time `gorm:"type:DATE;not null"`
	ThursdayOfBreak time.Time `gorm:"type:DATE;not null"`
}
