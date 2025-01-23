package login

import (
	"errors"
	"nova-hrm/app/domain/nova-hrm/models"
	"nova-hrm/app/domain/nova-hrm/utils"
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

func updateJwtToken(id string, username string, role string) (string, error) {
	accessToken := utils.GenerateJWT(username, role)

	error := updateUserAccessToken(id, accessToken)

	if error != nil {
		return "", error
	}

	return accessToken, nil
}
