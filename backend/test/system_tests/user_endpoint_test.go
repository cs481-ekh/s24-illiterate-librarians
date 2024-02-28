package system_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"LiteracyLink.com/backend/api/model"
	"LiteracyLink.com/backend/test/testInit"
	"github.com/google/uuid"
)

func TestUserEndpointTest(t *testing.T) {
	// Create a router from your main application
	router := testInit.InitRouterFunction()

	// Create a new UUID for the UserId
	userID := uuid.New()

	// Create a User object
	user := model.User{
		UserId:            userID.String(),
		UserType:          model.Client,
		FirstName:         "John",
		LastName:          "Doe",
		Email:             "john.doe@example.com",
		Address:           "123 Main Street",
		PhoneNumber:       "+1234567890",
		PrefContactMethod: "email",
	}

	// Create a payload map
	payload := map[string]interface{}{
		"user": user,
	}

	// Define test cases
	testCases := []testInit.TestCase{
		{Method: "POST", Endpoint: "/user/create", Payload: payload, ExpectedStatus: http.StatusCreated},
		{Method: "GET", Endpoint: "/user/login", ExpectedStatus: http.StatusOK},
		{Method: "GET", Endpoint: "/user/lookup", Payload: nil, ExpectedStatus: http.StatusOK},
		{Method: "GET", Endpoint: "/user/lookup/123", Payload: nil, ExpectedStatus: http.StatusOK},
		{Method: "PUT", Endpoint: "/user/update/123", ExpectedStatus: http.StatusAccepted},
		{Method: "PUT", Endpoint: "/user/update/password/123", ExpectedStatus: http.StatusAccepted},
		{Method: "DELETE", Endpoint: "/user/delete/123", Payload: nil, ExpectedStatus: http.StatusOK},
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
