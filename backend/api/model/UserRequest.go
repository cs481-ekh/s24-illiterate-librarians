package model

import "github.com/google/uuid"

type UserRequest struct {
	UserID uuid.UUID `json:"parent_id" binding:"required"`
}