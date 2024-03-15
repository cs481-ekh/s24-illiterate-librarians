package model

type Tutor struct {
	TutorID []byte `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"tutor_id"`
	UserID  []byte `gorm:"type:BINARY(16);NOT NULL" json:"user_id"`

	// Define a field to represent the relationship with the User model
	User User `gorm:"foreignKey:UserID"`
}

func (Tutor) TableName() string {
	return "Tutors" // Specify the exact table name in the database
}
