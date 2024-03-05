package unit_test

import (
	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/db"
	"testing"
)

func TestCreateUser(t *testing.T) {
	request := model.User{
		Username:       "testuser",
		PasswordHash:   "hashedpassword",
		Email:          "test@example.com",
		PhoneNumber:    "1234567890",
		FirstName:      "Test",
		LastName:       "User",
		PrefMethodComm: "Email",
		MailingAddress: "123 Main St, City",
	}

	connection := db.ConnectDB()
	err := db.CreateUser(request, connection)
	if err != nil {
		t.Errorf("fail: %s", err)
	}
	db.DisconnectDB(connection)
}
