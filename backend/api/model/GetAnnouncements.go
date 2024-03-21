package model

import "time"

type GetAnnouncements struct {
	AnnouncementID string    `json:"announcement_id"`
	AText          string    `json:"a_text"`
	CreatedDate    time.Time `json:"created_date"`
}

func (GetAnnouncements) TableName() string {
	return "Announcements"
}
