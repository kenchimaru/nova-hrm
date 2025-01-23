package register

import (
	"net/http"
	"nova-hrm/app/domain/nova-hrm/middleware"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all the application routes
func RegisterRoutes(router *mux.Router) {
	{
		router.Handle("/user/add", middleware.AuthAdmin(http.HandlerFunc(HandleAddUser))).Methods(http.MethodPost)
	}
}
