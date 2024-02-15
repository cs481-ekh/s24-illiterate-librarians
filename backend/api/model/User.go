package model

type User struct {
	UserId            string   `json:"userId"`
	UserType          UserType `json:"userType"`
	FirstName         string   `json:"firstName"`
	LastName          string   `json:"lastName"`
	Email             string   `json:"email"`
	Address           string   `json:"address"`
	PhoneNumber       string   `json:"phoneNumber"`
	PrefContactMethod string   `json:"prefContactMethod"`
}

type UserType string

const (
	Client     UserType = "client"
	Tutor      UserType = "tutor"
	Instructor UserType = "instructor"
)
