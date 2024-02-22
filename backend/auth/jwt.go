package auth

import (
	"LiteracyLink.com/backend/api/model"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("Literacy_Link_Secret_Key"))

// GenerateUserToken generates a new JWT token for the given user ID
func GenerateUserToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiration time 1 day
	})

	return token.SignedString(secretKey)
}

func GenerateUserJWT(user model.User) (string, error) {
	// Create a new token with custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.UserID,
		"username":  user.Username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours)
		"issued_at": time.Now().Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyUserToken verifies the JWT token and returns the userID if valid
func VerifyUserToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", jwt.ErrSignatureInvalid
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return "", jwt.ErrSignatureInvalid
	}

	return userID, nil
}
