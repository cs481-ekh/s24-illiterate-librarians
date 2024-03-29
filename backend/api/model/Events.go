package model

import "time"

type Events struct {
	EventID      []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"event_id"`
	EventTitle   string    `gorm:"type:VARCHAR(255); NOT NULL" json:"event_title"`
	EventDescrip string    `gorm:"type:VARCHAR(512); NOT NULL" json:"event_descrip"`
	DueDate      time.Time `gorm:"type:DATETIME;NOT NULL" json:"due_date"`
}

func (Events) TableName() string {
	return "Events"
}
