package model

import (
	"time"
)

type Semester struct {
	SemesterID      []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"semester_id"`
	SpringOrFall    string    `gorm:"type:VARCHAR(20);not null" json:"spring_or_fall"`
	FirstSchoolDay  time.Time `gorm:"type:DATE;not null" json:"first_school_day"`
	LastSchoolDay   time.Time `gorm:"type:DATE;not null" json:"last_school_day"`
	FirstTutorDay   time.Time `gorm:"type:DATE;not null" json:"first_tutor_day"`
	LastTutorDay    time.Time `gorm:"type:DATE;not null" json:"last_tutor_day"`
	TuesdayOfBreak  time.Time `gorm:"type:DATE;not null" json:"tuesday_of_break"`
	ThursdayOfBreak time.Time `gorm:"type:DATE;not null" json:"thursday_of_break"`
}
