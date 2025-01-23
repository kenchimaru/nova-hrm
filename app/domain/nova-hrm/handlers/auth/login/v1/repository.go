package login

import (
	"context"
	"errors"
	"nova-hrm/app/domain/nova-hrm/database"
	"nova-hrm/app/domain/nova-hrm/models"
	"nova-hrm/app/domain/nova-hrm/utils"
	"time"

	"cloud.google.com/go/firestore"
)

func getUserByUsername(username string) (*models.User, error) {
	ctx := context.Background()
	fc, err := database.GetFirestoreClient(ctx)

	if err != nil {
		return nil, err
	}

	user, err := fc.Collection("Users").Where("Username", "==", username).Limit(1).Documents(ctx).Next()
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

func updateUserAccessToken(user_id string, accessToken string) error {
	ctx := context.Background()
	fc, err := database.GetFirestoreClient(ctx)

	if err != nil {
		return err
	}

	err = fc.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		docRef := fc.Collection("Users").Doc(user_id)
		err := tx.Set(docRef, map[string]interface{}{
			"AccessToken": accessToken,
		}, firestore.MergeAll)

		if err != nil {
			return errors.New("error updating access token")
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func createAuthLog(user_id string, ipAddress string, action string, time time.Time) error {
	ctx := context.Background()
	fc, err := database.GetFirestoreClient(ctx)

	if err != nil {
		return err
	}

	docRef := fc.Collection("AuthLogs").NewDoc()

	err = fc.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {

		authLog := models.AuthLog{
			ID:        docRef.ID,
			UserID:    user_id,
			IPAddress: ipAddress,
			Action:    action,
			CreatedAt: time,
			UpdatedAt: time,
		}

		err := tx.Set(docRef, authLog)

		if err != nil {
			return errors.New("error creating auth log")
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
