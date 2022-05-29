package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPass(password string) string {
	pass := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}
