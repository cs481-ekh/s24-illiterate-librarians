package model

type TutorSessionLinker struct {
	SemesterTutoringID    []byte    `gorm:"type:BINARY(16);NOT NULL" json:"semester_tutoring_obj"`
	TutorSessionID    []byte    `gorm:"type:BINARY(16);NOT NULL" json:"tutor_session"`

	// Define fields for relationships with other tables
	SemesterTutObj    SemesterTutoringObj    `gorm:"foreignKey:SemesterTutoringID"`
	TutorSession TutorSession `gorm:"foreignKey:SemesterID"`
}
