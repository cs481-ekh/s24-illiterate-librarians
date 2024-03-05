package model

type ClientSessionRequest struct {
	TutorSessionID string `json:"tutor_session_id" binding:"required"`
}