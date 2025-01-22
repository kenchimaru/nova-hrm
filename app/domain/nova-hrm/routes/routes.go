package routes

import (
	"net/http"
	liveness "nova-hrm/app/domain/nova-hrm/handlers"
	"nova-hrm/app/domain/nova-hrm/handlers/auth/register/v1"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all the application routes
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/liveness", liveness.Liveness).Methods(http.MethodGet)

	api := router.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()
	{
		register.RegisterRoutes(v1)
	}
}
