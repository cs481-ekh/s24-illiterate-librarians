package model


type PassUpdate struct {
	UserID string `json:"user_id" binding:"required"` 
	PasswordHash string `json:"password" binding:"required"`
}