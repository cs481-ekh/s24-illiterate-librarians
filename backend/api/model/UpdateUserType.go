package model

type UpdateUserType struct {
	AdminID  string `json:"admin_id"`
	UserID   string `json:"user_id"`
	UserType string `json:"user_type"`
}
