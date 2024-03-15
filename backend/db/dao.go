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

func Login(request model.LoginRequest, db *gorm.DB) (model.User, string, error) {
	var user model.User
	result := db.Where("username = ?", request.Username).First(&user)
	if result.Error != nil {
		return user, "", result.Error
	}
	userType, _ := FindUserTable(user.UserID.String(), db)

	return user, userType, nil
}

func FindUserTable(userID string, db *gorm.DB) (string, error) {
	var parent model.Parent
	if err := db.Where("user_id = UUID_TO_BIN(?)", userID).First(&parent).Error; err == nil {
		return "parent", nil
	}

	var tutor model.Tutor
	if err := db.Where("user_id = UUID_TO_BIN(?)", userID).First(&tutor).Error; err == nil {
		return "tutor", nil
	}

	var instructor model.Instructor
	if err := db.Where("user_id = UUID_TO_BIN(?)", userID).First(&instructor).Error; err == nil {
		return "instructor", nil
	}

	return "user", gorm.ErrRecordNotFound
}

func CreateUser(request model.User, db *gorm.DB) error {
	result := db.Create(request)

	if result.Error != nil {
		return result.Error
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

func GetSessionsByUserIdAndType(userID string, userType string, db *gorm.DB) ([]model.TutorSession, error) {
	var sessions []model.TutorSession

	switch userType {
	case "parent":
		if err := db.Joins("JOIN Parents ON Tutor_session.parent_id = Parents.parent_id").
			Where("Parents.user_id = UUID_TO_BIN(?)", userID).
			Find(&sessions).Error; err != nil {
			return nil, err
		}
	case "tutor":
		err := db.Where("tutor_id = UUID_TO_BIN(?)", userID).Find(&sessions).Error
		if err != nil {
			return nil, err
		}
	default:
		return nil, gorm.ErrRecordNotFound
	}

	return sessions, nil
}

func GetEOSSurvey(request model.EOSRequest, db *gorm.DB) (model.EOSParentSurvey, error) {

	var EOS model.EOSParentSurvey
	result := db.Where("tutor_session_id = ?", request.EOSPSID).First(&EOS)
	if result.Error != nil {
		return EOS, result.Error
	}
	return EOS, nil
}
