package db

import (
	"LiteracyLink.com/backend/api/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("Literacy_Link_DB_USERNAME"),
		os.Getenv("Literacy_Link_DB_PASSWORD"),
		os.Getenv("Literacy_Link_DB_HOST"), // should be set to db
		os.Getenv("Literacy_Link_DB_PORT"),
		os.Getenv("Literacy_Link_DB_NAME"),
	)
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s \n with dsn %s", err, dsn)
	}
	log.Println("Connected to DB!")
	return db
}

func DisconnectDB(db *gorm.DB) {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Println("Error getting DB instance:", err)
		} else {
			err := sqlDB.Close()
			if err != nil {
				log.Println("Error closing database connection:", err)
			} else {
				log.Println("Database connection closed successfully.")
			}
		}
	}
}

func Login(request model.LoginRequest, db *gorm.DB) (model.User, error) {
	var user model.User
	result := db.Where("username = ?", request.Username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func CreateUser(request model.User, db *gorm.DB) error {
	result := db.Create(request)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePass(request model.PassUpdate, db *gorm.DB) error {
	var user model.User
	result := db.Save(request) //will need to test this
	if result.Error != nil {
		return user, result.Error
	}
	return nil
}

func SubmitApp(request model.TutoringApplication, db *gorm.DB) error {

	result := db.Create(request)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateApp(request model.TutoringApplication, db *gorm.DB) error {

	result := db.Save(request)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetApp(request model.AppRequest, db *gorm.DB) (model.TutoringApplication, error) {

	var app model.TutoringApplication
	result := db.Where("parent_id = ? AND child_id = ? AND desired_semester_id = ?", request.Parent, request.Child, request.Semester).First(&app)
	if result.Error != nil {
		return app, result.Error
	}
	return app, nil
}

func GetClientSession(request model.ClientSessionRequest, db *gorm.DB) (model.TutorSession, error) {

	var ses model.TutorSession
	result := db.Where("tutor_session_id = ?", request.TutorSessionID).First(&ses)
	if result.Error != nil {
		return ses, result.Error
	}
	return ses, nil
}

func GetEOSSurvey(request model.EOSRequest, db *gorm.DB) (model.EOSParentSurvey, error) {

	var EOS model.EOSParentSurvey
	result := db.Where("tutor_session_id = ?", request.EOSPSID).First(&EOS)
	if result.Error != nil {
		return EOS, result.Error
	}
	return EOS, nil
}

func SubmitEOSSurvey(request model.EOSParentSurvey, db *gorm.DB) error {

	result := db.Create(request)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
