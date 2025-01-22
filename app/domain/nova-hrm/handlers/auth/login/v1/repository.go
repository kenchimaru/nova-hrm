package login

import (
	"context"
	"errors"
	"nova-hrm/app/domain/nova-hrm/database"

	"cloud.google.com/go/firestore"
)

func GetUserByUsername(username string) (*firestore.DocumentSnapshot, error) {
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

	return user, nil
}
