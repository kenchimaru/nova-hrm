package middleware

import (
	"context"
	"errors"
	"nova-hrm/app/domain/nova-hrm/database"
	"nova-hrm/app/domain/nova-hrm/models"
	"nova-hrm/app/domain/nova-hrm/utils"
)

func getUserByID(ID string) (*models.User, error) {
	ctx := context.Background()
	fc, err := database.GetFirestoreClient(ctx)

	if err != nil {
		return nil, err
	}

	user, err := fc.Collection("Users").Doc(ID).Get(ctx)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	var userModel models.User
	err = utils.MapDocumentToModel(user, &userModel)

	if err != nil {
		return nil, err

	}
	return &userModel, nil
}
