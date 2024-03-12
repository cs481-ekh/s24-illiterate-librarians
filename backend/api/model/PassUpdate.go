package model

type PassUpdate struct {
	Username string `json:"username" binding:"required"` //is binding required for these two?
	Password string `json:"password" binding:"required"`
}