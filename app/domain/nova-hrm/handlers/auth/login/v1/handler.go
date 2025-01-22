package login

import (
	"encoding/json"
	"log"
	"net/http"
	"nova-hrm/app/domain/nova-hrm/response"
	"nova-hrm/app/domain/nova-hrm/validators"
)

// LoginRequest represents the structure of a login request payload
type LoginRequest struct {
	Username string `json:"username" validate:"required,min=6,max=50,startswithalpha"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

// Login handles user login with mux router
func HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest

	// Parse JSON request
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		response.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	// Validate the request
	if err := validators.Validate.Struct(loginReq); err != nil {
		log.Printf("Validation error: %v", err)
		response.Error(w, "Bad request", http.StatusBadRequest)

		return
	}

	// Authenticate user
	err := AuthenticateWithPassword(loginReq.Username, loginReq.Password)

	if err != nil {
		response.Error(w, "Invalid username or password", http.StatusUnauthorized)

		return
	}

	response.Success(w, "Login successful", http.StatusOK, nil)
}
