package helper

import (
	"log"

	"github.com/golang-jwt/jwt"
)

var mySecretKey = []byte("$sussybaka")

func GenerateJWT(id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["userid"] = id

	tokenString, err := token.SignedString(mySecretKey)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}
