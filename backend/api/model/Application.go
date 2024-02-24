package model

import (
	"gorm.io/gorm"
)

// NOTE: not all of the things that will be asked in the application will
// be stored here / in the app_for_tutoring table. Some things like the days
// that parents will be unavailable will be stored in other tables. For a
// full list of questions in the application, look at the original applicaiton.

// TODO: list out in the comments here the questions that should be asked for the benefit of future developers

type Application struct {
	gorm.Model
	// ParentId  string `json:"parentId"`
	// ChildId   string `json:"childId"`
	// FirstName string `json:"firstName"`
	// LastName  string `json:"lastName"`
	// Grade     int8   `json:"grade"`
}
