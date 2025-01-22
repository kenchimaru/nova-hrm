package login

import (
	"errors"

	"cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
)

// CheckPassword compares a hashed password with its possible plaintext equivalent
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}

func AuthenticateWithPassword(username string, password string) error {

	user, err := GetUserByUsername(username)

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	passwordData, err := user.DataAtPath(firestore.FieldPath{"Password"})
	if err != nil {
		return errors.New("error retrieving password data")
	}

	userPassword, ok := passwordData.(string)
	if !ok {
		return errors.New("password data is not a string")
	}

	if !CheckPassword(userPassword, password) {
		return errors.New("invalid password")
	}

	return nil
}
