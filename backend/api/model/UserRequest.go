package model

import "github.com/google/uuid"

type UserRequest struct {
	UserID uuid.UUID `json:"user_id" binding:"required"`
}