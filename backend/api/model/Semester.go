package model

type Semester struct {
	Year      int    `json:"year"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Term      string `json:"term"`
}
