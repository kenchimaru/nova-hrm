package register

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(
	username string,
	encPassword string,
	email string,
) (string, error) {
	id, err := CreateUserWithPassword(
		username,
		encPassword,
		email,
	)

	return id, err
}

func HashPassword(password string) (string, error) {
	log.Printf("Hashing password %v", password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
