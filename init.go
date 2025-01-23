package main

import (
	"context"
	"errors"
	"fmt"
	"nova-hrm/app/domain/nova-hrm/database"
	"nova-hrm/app/domain/nova-hrm/models"
	"nova-hrm/app/domain/nova-hrm/utils"
	"strings"

	"cloud.google.com/go/firestore"
)

func init() {
	encPassword, err := utils.HashPassword("secret")
	if err != nil {
		panic(err)
	}

	user := models.User{
		Username: "admin",
		Email:    "admin@admin.com",
		Password: encPassword,
		Role:     "admin",
	}

	id, err := createUserWithPassword(
		user.Username,
		user.Password,
		user.Email,
		user.Role,
	)

	if err != nil {
		fmt.Printf("Error creating admin user: %v\n", err)
	}

	if id != "" {
		fmt.Printf("Admin user created with ID: %v\n", id)

	}
}

// CreateUser creates a new user in the database
func createUserWithPassword(
	username string,
	password string,
	email string,
	role string,
) (string, error) {
	ctx := context.Background()
	fc, err := database.GetFirestoreClient(ctx)

	if err != nil {
		return "", err
	}

	d := fc.Collection("Users").NewDoc()
	// Convert the struct to a map dynamically
	user := models.User{
		ID:       d.ID,
		Username: username,
		Email:    email,
		Password: password,
		Role:     role,
	}

	err = fc.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		emailRef := fc.Collection("Emails").Doc(email)
		emailSanpShot, err := tx.Get(emailRef)

		if err != nil && !strings.Contains(err.Error(), "not found") {
			return errors.New("bad request")
		}

		if emailSanpShot.Exists() {
			return errors.New("admin already exist")
		}

		usernameRef := fc.Collection("Usernames").Doc(username)
		usernameSnapShot, err := tx.Get(usernameRef)

		if err != nil && !strings.Contains(err.Error(), "not found") {
			return errors.New("bad request")
		}

		if usernameSnapShot.Exists() {
			return errors.New("admin already exist")
		}

		tx.Create(d, user)

		refPath := "Users" + "/" + d.ID
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
