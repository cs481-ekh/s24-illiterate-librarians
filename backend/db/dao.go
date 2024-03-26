package db

import (
	"LiteracyLink.com/backend/api/model"
	"errors"
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
	userType, _ := FindUserTableStringKey(user.UserID.String(), db)

	return user, userType, nil
}

func FindUserTableStringKey(userID string, db *gorm.DB) (string, error) {
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

	return "user", nil
}

func SetUserType(UpdatedUserType model.UpdateUserType, db *gorm.DB) error {
	// Check if the admin has permissions
	admin, err := FindUserTableStringKey(UpdatedUserType.AdminID, db)
	if err != nil {
		return err
	}
	if admin != "instructor" {
		return errors.New("lacks permissions")
	}

	// Begin a transaction
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Delete old rows from other tables
	if err := tx.Exec("DELETE FROM Instructors WHERE user_id = UUID_TO_BIN(?)", UpdatedUserType.UserID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Exec("DELETE FROM Tutors WHERE user_id = UUID_TO_BIN(?)", UpdatedUserType.UserID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Exec("DELETE FROM Parents WHERE user_id = UUID_TO_BIN(?)", UpdatedUserType.UserID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Insert the user into the new user type table
	switch UpdatedUserType.UserType {
	case "parent":
		if err := tx.Exec("INSERT INTO Parents (user_id) VALUES (UUID_TO_BIN(?))", UpdatedUserType.UserID).Error; err != nil {
			tx.Rollback()
			return err
		}
	case "instructor":
		if err := tx.Exec("INSERT INTO Instructors (user_id) VALUES (UUID_TO_BIN(?))", UpdatedUserType.UserID).Error; err != nil {
			tx.Rollback()
			return err
		}
	case "tutor":
		if err := tx.Exec("INSERT INTO Tutors (user_id) VALUES (UUID_TO_BIN(?))", UpdatedUserType.UserID).Error; err != nil {
			tx.Rollback()
			return err
		}
	default:
		tx.Rollback()
		return gorm.ErrInvalidData
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func CreateUser(request model.User, db *gorm.DB) error {
	// Begin a transaction
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Create the user entry
	if err := tx.Create(&request).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Retrieve the generated UserID
	var userID string
	if err := tx.Raw("SELECT BIN_TO_UUID(user_id) FROM Users WHERE username = ?;", request.Username).Scan(&userID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Exec("INSERT INTO Parents (user_id) VALUES (UUID_TO_BIN(?));", userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return err
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

func GetTutorsSessions(userID string, db *gorm.DB) ([]model.GetTutorSessionReply, error) {
	var sessions []model.GetTutorSessionReply
	err := db.Table("Tutor_session AS ts").
		Select("BIN_TO_UUID(ts.tutor_session_id) as tutor_session_id, ts.zoom_join_link, ts.zoom_recording_link, ts.meeting_date, ts.parent_avail").
		Joins("JOIN Tutor_linker tl ON ts.semester_tutoring_obj = tl.semester_tutoring_obj").
		Joins("JOIN Tutors t ON tl.tutor_id = t.tutor_id").
		Joins("JOIN Users AS u ON t.user_id = u.user_id").
		Where("u.user_id = UUID_TO_BIN(?)", userID).
		Find(&sessions).Error
	return sessions, err
}

func GetClientsSessions(userID string, db *gorm.DB) ([]model.GetTutorSessionReply, error) {
	var sessions []model.GetTutorSessionReply
	err := db.Table("Tutor_session AS ts").
		Select("BIN_TO_UUID(ts.tutor_session_id) as tutor_session_id, ts.zoom_join_link, ts.zoom_recording_link, ts.meeting_date, ts.parent_avail").
		Joins("JOIN Semester_tutoring_obj sto ON ts.semester_tutoring_obj = sto.semester_tutoring_id").
		Joins("JOIN Parents p ON sto.parent_id = p.parent_id").
		Joins("JOIN Users u ON p.user_id = u.user_id").
		Where("u.user_id = UUID_TO_BIN(?)", userID).
		Find(&sessions).Error
	return sessions, err
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

func GetAnnouncements(db *gorm.DB) ([]model.GetAnnouncements, error) {
	var result []model.GetAnnouncements
	err := db.
		Select("BIN_TO_UUID(announcement_id) as announcement_id, a_text, created_date").
		Where("created_date between date_sub(now(),INTERVAL 2 WEEK) and now();").Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetEvents(db *gorm.DB) ([]model.GetEvents, error) {
	var result []model.GetEvents
	err := db.
		Select("BIN_TO_UUID(event_id) as event_id, event_title, event_descrip, created_date, due_date").
		Where("due_date >= date_sub(NOW(), INTERVAL 1 WEEK)").Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetChildren(id string, db *gorm.DB) ([]model.ChildJSON, error) {
	var result []model.ChildJSON
	err := db.
		Table("Users").
		Select("BIN_TO_UUID(c.child_id) as child_id, BIN_TO_UUID(p.parent_id) as parent_id, c.grade, c.first_name, c.last_name").
		Joins("JOIN Parents p ON Users.user_id = p.user_id").
		Joins("JOIN Child c ON p.parent_id = c.parent_id").
		Where("Users.user_id = UUID_TO_BIN(?)", id).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
