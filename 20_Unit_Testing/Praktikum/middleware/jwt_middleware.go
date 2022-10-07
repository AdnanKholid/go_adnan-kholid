package middleware

import (
	"echo-api/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_KEY))
}
