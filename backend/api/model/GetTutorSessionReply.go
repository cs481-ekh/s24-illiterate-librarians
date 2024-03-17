package model

import "time"

type GetTutorSessionReply struct {
	TutorSessionID    string    `json:"tutor_session_id"`
	ZoomJoinLink      string    `json:"zoom_join_link"`
	ZoomRecordingLink string    `json:"zoom_recording_link"`
	MeetingDate       time.Time `json:"meeting_date"`
	ParentAvail       bool      `json:"parent_avail"`
	SemesterType      string    `json:"semester_type"`
	SemesterStart     time.Time `json:"semester_start"`
	SemesterEnd       time.Time `json:"semester_end"`
}
