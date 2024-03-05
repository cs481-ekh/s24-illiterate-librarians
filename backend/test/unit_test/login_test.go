package unit_test

import (
	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	request := model.LoginRequest{
		Username: "john_doe",
		Password: "hashed_password_1",
	}

	connection := db.ConnectDB()
	user, err := db.Login(request, connection)
	if err != nil {
		t.Errorf("fail: %s", err)
	}
	db.DisconnectDB(connection)
	fmt.Println(user.Username)
	fmt.Println(user.UserID.String())
}
