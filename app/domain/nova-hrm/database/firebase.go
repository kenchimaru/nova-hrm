package database

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// FirestoreClient wraps the Firestore client to ensure it is closed after use
var Client *firestore.Client

// NewFirestoreClient initializes a new Firestore client
func InitFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	if Client != nil {
		return Client, nil
	}

	log.Println("Initializing Firestore client")
	sa := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	Client = client

	return client, nil
}

func GetFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	if Client == nil {
		fmt.Println("Client is nil")
		return InitFirestoreClient(ctx)
	}

	return Client, nil
}

// Close closes the Firestore client
func Close(fc *firestore.Client) {
	log.Println("Closing Firestore client")
	if fc != nil {
		fc.Close()
	}
}
