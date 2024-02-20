package system_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSurveyEndpoints(t *testing.T) {
	// Create a router from your main application
	router := InitRouterFunction()

	// Define test cases
	testCases := []testCase{
		{Method: "POST", Endpoint: "/survey/pre_semester_survey/123", Payload: map[string]string{"answer": "example"}, ExpectedStatus: http.StatusAccepted},
		{Method: "POST", Endpoint: "/survey/after_semester_survey/123", Payload: map[string]string{"answer": "example"}, ExpectedStatus: http.StatusAccepted},
		{Method: "GET", Endpoint: "/survey/pre_semester_survey/taken/123", ExpectedStatus: http.StatusOK},
		{Method: "GET", Endpoint: "/survey/after_semester_survey/taken/123", ExpectedStatus: http.StatusOK},
		{Method: "GET", Endpoint: "/survey/after_semester_survey/123", ExpectedStatus: http.StatusOK},
		{Method: "GET", Endpoint: "/survey/pre_semester_survey/123", ExpectedStatus: http.StatusOK},

		// Add more test cases as needed
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
