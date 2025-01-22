package register

import (
	"context"
	"errors"
	"log"
	"nova-hrm/app/domain/nova-hrm/database"
	"nova-hrm/app/domain/nova-hrm/models"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"
)

var collection = "users"

// CreateUser creates a new user in the database
func CreateUser(user models.User) (*firestore.DocumentRef, error) {
	ctx := context.Background()
	fc, err := database.GetFirestoreClient(ctx)

	if err != nil {
		return nil, err
	}

	d := fc.Collection(collection).NewDoc()
	log.Printf("Document ID: %s", d.ID)
	// Convert the struct to a map dynamically
	dataMap := make(map[string]interface{})
	dataValue := reflect.ValueOf(user)
	dataType := reflect.TypeOf(user)

	// Populate the map with the struct fields
	for i := 0; i < dataValue.NumField(); i++ {
		fieldName := dataType.Field(i).Name
		fieldValue := dataValue.Field(i).Interface()
		dataMap[fieldName] = fieldValue
	}

	dataMap["ID"] = d.ID

	err = fc.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		emailRef := fc.Collection("emails").Doc(user.Email)
		emailSanpShot, err := tx.Get(emailRef)

		if err != nil && !strings.Contains(err.Error(), "not found") {
			return errors.New("bad_request")
		}

		log.Printf("Email snapshot: %v", emailSanpShot.Exists())
		if emailSanpShot.Exists() {
			return errors.New("email already exist")
		}

		usernameRef := fc.Collection("usernames").Doc(user.Username)
		usernameSnapShot, err := tx.Get(usernameRef)

		if err != nil && !strings.Contains(err.Error(), "not found") {
			return errors.New("bad_request")
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
		return nil, err
	}

	return d, nil
}
