package middlewares

import (
	"acp8/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userId int) (string, error) {

	// payload
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour

	// header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// kunci
	return token.SignedString([]byte(constants.SECRET_JWT))
}
