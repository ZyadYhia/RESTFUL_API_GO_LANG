package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "supersecret"

func GenerateJWTToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//Data will be included in token
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	// secret key will be used to verify the token later
	return token.SignedString([]byte(secretKey))
}
func VerifyJWTToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// type check syntax if value stored in token.Method
		//is same like (*jwt.SigningMethodHMAC)
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected Signing Method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("Couldn't Parse Token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("Invalid Token")
	}
	/*
		we can use next lines if we wanted to extract data from token
	*/

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid Token Claims")
	}
	//	email := claims["email"].(string)
	// the interface returned float64 and we have to convert to int64
	userId := int64(claims["userId"].(float64))

	return userId, nil
}
