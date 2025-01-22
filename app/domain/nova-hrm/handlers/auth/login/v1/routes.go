package login

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all the application routes
func RegisterRoutes(router *mux.Router) {
	{
		router.HandleFunc("/login", HandleUserLogin).Methods(http.MethodPost)
	}
}
