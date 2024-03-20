package model

type TutorLinker struct {
	SemesterTutoringID    []byte    `gorm:"type:BINARY(16);NOT NULL" json:"semester_tutoring_obj"`
	TutorID    []byte    `gorm:"type:BINARY(16);NOT NULL" json:"tutor_id"`

	// Define fields for relationships with other tables
	SemesterTutObj    SemesterTutoringObj    `gorm:"foreignKey:SemesterTutoringID"`
	Tutor Tutor `gorm:"foreignKey:TutorID"`
}
