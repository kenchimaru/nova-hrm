package register

import (
	"context"
	"errors"
	"log"
	"nova-hrm/app/domain/nova-hrm/database"
	"nova-hrm/app/domain/nova-hrm/models"
	"strings"

	"cloud.google.com/go/firestore"
)

var collection = "Users"

// CreateUser creates a new user in the database
func CreateUserWithPassword(
	username string,
	password string,
	email string,
) (string, error) {
	ctx := context.Background()
	fc, err := database.GetFirestoreClient(ctx)

	if err != nil {
		return "", err
	}

	d := fc.Collection(collection).NewDoc()
	// Convert the struct to a map dynamically
	dataMap := make(map[string]interface{})

	dataMap["ID"] = d.ID
	dataMap["Username"] = username
	dataMap["Email"] = email
	dataMap["Password"] = password

	err = fc.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		emailRef := fc.Collection("Emails").Doc(email)
		emailSanpShot, err := tx.Get(emailRef)

		if err != nil && !strings.Contains(err.Error(), "not found") {
			return errors.New("bad request")
		}

		log.Printf("Email snapshot: %v", emailSanpShot.Exists())
		if emailSanpShot.Exists() {
			return errors.New("email already exist")
		}

		usernameRef := fc.Collection("Usernames").Doc(username)
		usernameSnapShot, err := tx.Get(usernameRef)

		if err != nil && !strings.Contains(err.Error(), "not found") {
			return errors.New("bad request")
		}

		if usernameSnapShot.Exists() {
			return errors.New("username already exist")
		}

		tx.Create(d, dataMap)

		refPath := collection + "/" + d.ID
		docIndex := models.DocumentIndex{
			RefPath: refPath,
		}

		tx.Create(emailRef, docIndex)
		tx.Create(usernameRef, docIndex)

		return nil
	})

	if err != nil {
		return "", err
	}

	return d.ID, nil
}
