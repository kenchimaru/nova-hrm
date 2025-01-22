package main

import (
	"context"
	"log"
	"net/http"
	"nova-hrm/app/config"
	"nova-hrm/app/domain/nova-hrm/database"
	"nova-hrm/app/domain/nova-hrm/routes"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	ctx := context.Background()

	log.Println("Starting server")
	firebaseClient, err := database.InitFirestoreClient(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v", err)
	}

	router := mux.NewRouter()

	defer database.Close(firebaseClient)

	router.Use(bodyParserMiddleware)

	routes.RegisterRoutes(router)

	url := cfg.Host + ":" + cfg.Port
	http.ListenAndServe(url, router)
}
