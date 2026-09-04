package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtSecret = "onma-cycle-secret-key"

func GenerateToken(userID int, email string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {

			// Make sure the token was signed using HMAC.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(jwtSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

// package utils

// import (
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// const jwtSecret = "onma-cycle-secret-key"

// func GenerateToken(userID int, email string) (string, error) {

// 	claims := jwt.MapClaims{
// 		"user_id": userID,
// 		"email":   email,
// 		"exp":     time.Now().Add(24 * time.Hour).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	return token.SignedString([]byte(jwtSecret))
// }
