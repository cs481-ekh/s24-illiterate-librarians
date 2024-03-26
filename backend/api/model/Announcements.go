package model

import "time"

type Announcements struct {
	AnnouncementID []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"announcement_id"`
	AText          string    `gorm:"type:VARCHAR(512); NOT NULL" json:"a_text"`
	CreatedDate    time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_date"`
}

func (Announcements) TableName() string {
	return "Announcements"
}
