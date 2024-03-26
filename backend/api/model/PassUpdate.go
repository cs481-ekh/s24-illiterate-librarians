package model

import "github.com/google/uuid"

type PassUpdate struct {
	UserID uuid.UUID `json:"user_id" binding:"required"` 
	PasswordHash string `json:"password" binding:"required"`
}