package model

type UserRequest struct {
	UserID string `json:"user_id" binding:"required"`
}