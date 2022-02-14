package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userId int) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("InvalidToken")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
