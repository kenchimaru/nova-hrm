package login

import (
	"errors"
	"nova-hrm/app/domain/nova-hrm/models"
	"nova-hrm/app/domain/nova-hrm/utils"
	"time"
)

func authenticateWithPassword(username string, password string) (*models.User, error) {

	user, err := getUserByUsername(username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	passwordData := user.Password
	if err != nil {
		return nil, errors.New("error retrieving password data")
	}

	if !utils.CheckPassword(passwordData, password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func updateJwtToken(user_id string, username string, role string) (string, error) {
	accessToken := utils.GenerateJWT(user_id, username, role)

	error := updateUserAccessToken(user_id, accessToken)

	if error != nil {
		return "", error
	}

	return accessToken, nil
}

func addAuthLog(user_id string, ipAddress string, action string, time time.Time) error {
	err := createAuthLog(user_id, ipAddress, action, time)

	if err != nil {
		return err
	}

	return nil
}
