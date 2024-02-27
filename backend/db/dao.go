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
		os.Getenv("Literacy_Link_DB_HOST"),
		os.Getenv("Literacy_Link_DB_PORT"),
		os.Getenv("Literacy_Link_DB_NAME"),
	)
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
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
	result := db.Where("username = ? AND password_hash = ?", request.Username, request.Password).First(&user)
	if result.Error != nil {
		return user, errors.New("Invalid username or password")
	}
	return user, nil
}

// func SubmitApp(request model.Application, db *gorm.DB) (model.Application, error) {
// 	var app model.Application
// 	//Not 100% sure on what the below line should look like when trying to create a entry in the table vs a query of the table
// 	//result := db.Where("child_id = ? AND parent_id = ?", request.Username, request.Password).First(&user)
// 	if result.Error != nil {
// 		return app, errors.New("Invalid application submission")
// 	}
// 	return app, nil
// }
