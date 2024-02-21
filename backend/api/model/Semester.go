package model

import (
	"gorm.io/gorm"
)

type Semester struct {
	gorm.Model
	Year      int    `json:"year"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Term      string `json:"term"`
}
