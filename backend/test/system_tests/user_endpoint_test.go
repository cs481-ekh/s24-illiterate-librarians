package system_tests

import (
	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/auth"
	"LiteracyLink.com/backend/test/testInit"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func TestUserEndpointTest(t *testing.T) {
	// Create a router from your main application
	router := testInit.InitRouterFunction()

	// Create a new UUID for the UserId
	userID := uuid.New()

	// Create a User object
	user := model.User{
		UserID:         userID,
		Username:       "John_Doe",
		PasswordHash:   "Hashed_Password",
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john.doe@example.com",
		MailingAddress: "123 Main Street",
		PhoneNumber:    "+1234567890",
	}

	token, err := auth.GenerateJWT(userID)
	if err != nil {
		t.Errorf("Failed to generate JWT token for testing")
	}

	// Create a payload map
	payload := map[string]interface{}{
		"user": user,
	}

	header := map[string]interface{}{
		"Authorization": token,
	}

	// Define test cases
	testCases := []testInit.TestCase{
		{Method: "POST", Endpoint: "/user/create", Payload: payload, ExpectedStatus: http.StatusCreated},
		{Method: "GET", Endpoint: "/user/login", ExpectedStatus: http.StatusOK},
		{Method: "GET", Endpoint: "/user/lookup", Headers: header, ExpectedStatus: http.StatusOK},
		{Method: "GET", Endpoint: "/user/lookup/123", Headers: header, ExpectedStatus: http.StatusOK},
		{Method: "PUT", Endpoint: "/user/update/123", Headers: header, ExpectedStatus: http.StatusAccepted},
		{Method: "PUT", Endpoint: "/user/update/password/123", Headers: header, ExpectedStatus: http.StatusAccepted},
		{Method: "DELETE", Endpoint: "/user/delete/123", Headers: header, ExpectedStatus: http.StatusOK},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		t.Run(tc.Method+" "+tc.Endpoint, func(t *testing.T) {
			// Convert payload to JSON
			payloadBytes, err := json.Marshal(tc.Payload)
			if err != nil {
				t.Fatalf("Failed to marshal payload: %v", err)
			}

			// Create a request
			req, err := http.NewRequest(tc.Method, tc.Endpoint, bytes.NewBuffer(payloadBytes))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			// Set headers
			req.Header.Set("Authorization", token)

			// Create a response recorder
			rec := httptest.NewRecorder()

			// Serve the request using the main application's router
			router.ServeHTTP(rec, req)

			// Check the status code
			if rec.Code != tc.ExpectedStatus {
				t.Errorf("Expected status %d, got %d", tc.ExpectedStatus, rec.Code)
			}
		})
	}
}
