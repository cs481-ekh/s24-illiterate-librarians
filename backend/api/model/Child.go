package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Child struct {
	ChildID   []byte     `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"child_id"`
	ParentID  []byte     `gorm:"type:BINARY(16);NOT NULL" json:"parent_id"`
	BirthDate CustomTime `gorm:"type:DATE;NOT NULL" json:"birth_date" time_format:"RFC3339"`
	Grade     int8       `gorm:"type:TINYINT;NOT NULL" json:"grade"`
	FirstName string     `gorm:"type:VARCHAR(50);NOT NULL" json:"first_name"`
	LastName  string     `gorm:"type:VARCHAR(50);NOT NULL" json:"last_name"`
	CreatedAt time.Time  `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at"`

	// Define a field to represent the relationship with the Parent model
	Parent Parent `gorm:"foreignKey:ParentID"`
}

type CustomTime struct {
	time.Time
}

// UnmarshalJSON implements the Unmarshaler interface for CustomTime
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	str := string(b)
	str = str[1 : len(str)-1] // Trim quotes
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

// Value implements the driver Valuer interface
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

// Scan implements the Scanner interface
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Time = time.Time{}
		return nil
	}
	parsedTime, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("Unable to convert value to time.Time")
	}
	ct.Time = parsedTime
	return nil
}

type ChildJSON struct {
	ChildID   string `json:"child_id"`
	ParentID  string `json:"parent_id"`
	Grade     int8   `json:"grade"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (ChildJSON) TableName() string {
	return "Child" // Specify the exact table name in the database
}

func (Child) TableName() string {
	return "Child" // Specify the exact table name in the database
}
