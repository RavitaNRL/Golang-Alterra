package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"section_22/constants"
)

func CreateToken(userId int, name string) (string, error) {

	claims := jwt.MapClaims{}
	claims["userID"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))

}
