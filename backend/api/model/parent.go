package model

type Parent struct {
	ParentID []byte `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"parent_id"`
	UserID   []byte `gorm:"type:BINARY(16);not null" json:"user_id"`

	// Define a field to represent the relationship with the User model
	User User `gorm:"foreignKey:UserID"`
}
