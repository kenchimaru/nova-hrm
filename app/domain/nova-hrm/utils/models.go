package utils

import (
	"encoding/json"

	"cloud.google.com/go/firestore"
)

func MapDocumentToModel(doc *firestore.DocumentSnapshot, model interface{}) error {
	// Get the document data as a map
	data := doc.Data()

	// Convert the map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into the provided model
	err = json.Unmarshal(jsonData, model)
	if err != nil {
		return err
	}

	return nil
}
