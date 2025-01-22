package register

import "golang.org/x/crypto/bcrypt"

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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with its possible plaintext equivalent
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
