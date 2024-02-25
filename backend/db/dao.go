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

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
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
	return db
}

func disconnectDB(db *gorm.DB) {
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

func login(request model.LoginRequest) (model.User, error) {
	db := connectDB()
	var user model.User
	result := db.Where("username = ? AND password = ?", request.Username, request.Password).First(user)
	disconnectDB(db)
	if result.Error != nil {
		return user, errors.New("Invalid username or password")
	}
	return user, nil
}
