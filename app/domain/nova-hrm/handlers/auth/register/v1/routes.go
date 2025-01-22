package register

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all the application routes
func RegisterRoutes(router *mux.Router) {
	{
		router.HandleFunc("/register", HandleRegister).Methods(http.MethodPost)
	}
}
