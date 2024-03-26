package model

import "time"

type GetEvents struct {
	EventID      string    `json:"event_id"`
	EventTitle   string    `json:"event_title"`
	EventDescrip string    `json:"event_descrip"`
	DueDate      time.Time `json:"due_date"`
}

func (GetEvents) TableName() string {
	return "Events"
}
