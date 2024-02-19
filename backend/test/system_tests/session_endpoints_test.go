package system_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSessionEndpointTest(t *testing.T) {
	// Create a router from your main application
	router := InitRouterFunction()

	// Define test cases
	var sessionTestCases = []testCase{
		{Method: "GET", Endpoint: "/session/client/123", Payload: nil, ExpectedStatus: http.StatusOK},
		//{Method: "GET", Endpoint: "/session/client/invalid_user", Payload: nil, ExpectedStatus: http.StatusNotFound},
		{Method: "GET", Endpoint: "/session/tutor/456", Payload: nil, ExpectedStatus: http.StatusOK},
		//{Method: "GET", Endpoint: "/session/tutor/invalid_user", Payload: nil, ExpectedStatus: http.StatusNotFound},
		{Method: "GET", Endpoint: "/session/789", Payload: nil, ExpectedStatus: http.StatusOK},
		//{Method: "GET", Endpoint: "/session/invalid_session", Payload: nil, ExpectedStatus: http.StatusNotFound},
	}

	// Iterate through test cases
	for _, tc := range sessionTestCases {
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
