package model

type PassUpdate struct {
	UserID string `json:"user_id" binding:"required"` //is binding required for these two?
	PasswordHash string `json:"password" binding:"required"`
}