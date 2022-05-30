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

func ExtractClaims(at string) (jwt.MapClaims, bool) {
	hmacSecretString := mySecretKey
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(at, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
