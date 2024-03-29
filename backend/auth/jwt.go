package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("Literacy_Link_Secret_Key"))

// GenerateJWT generates a new JWT token for the given user ID
func GenerateJWT(userID uuid.UUID, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userID.String(),
		"userType": userType,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time 1 day
	})

	return token.SignedString(secretKey)
}

// VerifyJWT verifies the JWT token and returns the userID if valid
func VerifyJWT(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", jwt.ErrSignatureInvalid
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return "", "", jwt.ErrSignatureInvalid
	}
	userType, ok := claims["userType"].(string)
	if !ok {
		return "", "", jwt.ErrSignatureInvalid
	}
	_, err = uuid.Parse(userID)
	if err != nil {
		return "", "", err
	}
	return userID, userType, nil
}
