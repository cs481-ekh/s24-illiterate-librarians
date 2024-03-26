package model

import "time"

type TutoringSessionResponse struct {
	SessionId []byte    `json:"sessionId"`
	Title     string    `json:"title"`
	When      time.Time `json:"date"`
	ZoomLink  string    `json:"zoom_link"`
}
