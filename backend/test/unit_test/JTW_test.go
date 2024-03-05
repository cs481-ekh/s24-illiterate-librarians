package unit_test

import (
	"LiteracyLink.com/backend/auth"
	"LiteracyLink.com/backend/test/testInit"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJWTGenerationAndAuthentication(t *testing.T) {
	// Set up a test user ID (you can use a UUID library for dynamic IDs in production)
	testUserID, err := uuid.NewUUID()
	if err != nil {
		t.Errorf("Error generating uuid to test with")
	}

	// Generate a JWT for the test user
	token, err := auth.GenerateJWT(testUserID)
	if err != nil {
		t.Fatalf("Error generating JWT: %v", err)
	}

	// Create a test request with the generated token in the Authorization header
	req := httptest.NewRequest("GET", "/health", nil)
	req.Header.Set("Authorization", token)

	// Create a test response recorder to capture the response
	w := httptest.NewRecorder()

	// Initialize the router
	router := testInit.InitRouterFunction()

	// Serve the test request using the router
	router.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestJWTGenerationAndAuthenticationWithBadToken(t *testing.T) {
	// Set up a test case with a bad token
	badToken := "bad-token"

	// Create a new test request with the bad token in the Authorization header
	reqBadToken := httptest.NewRequest("GET", "/health", nil)
	reqBadToken.Header.Set("Authorization", badToken)

	// Create a new test response recorder
	wBadToken := httptest.NewRecorder()

	// Initialize the router
	router := testInit.InitRouterFunction()

	// Serve the request with the bad token
	router.ServeHTTP(wBadToken, reqBadToken)

	// Check the response status code (should be unauthorized)
	if wBadToken.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, wBadToken.Code)
	}
}
