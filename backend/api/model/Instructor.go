package model

type Instructor struct {
	InstructorID []byte `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"instructor_id"`
	UserID       []byte `gorm:"type:BINARY(16);NOT NULL" json:"user_id"`

	// Define a field to represent the relationship with the User model
	User User `gorm:"foreignKey:UserID"`
}

func (Instructor) TableName() string {
	return "Instructors" // Specify the exact table name in the database
}
