package register

import (
	"encoding/json"
	"net/http"
	"nova-hrm/app/domain/nova-hrm/models"
	"nova-hrm/app/domain/nova-hrm/response"
	"nova-hrm/app/domain/nova-hrm/validators"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=6,max=20,startswithalpha"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Email    string `json:"email" validate:"required,email"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var registerReq RegisterRequest

	// Parse JSON request
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		response.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	// Validate the request
	if err := validators.Validate.Struct(registerReq); err != nil {
		response.Error(w, "Bad request", http.StatusBadRequest)

		return
	}

	// Prepare user object
	user := models.User{
		Username: registerReq.Username,
		Password: registerReq.Password,
		Email:    registerReq.Email,
	}

	// Save user to Firestore
	docRef, err := CreateUser(user)
	if err != nil {
		response.Error(w, "Error creating user", http.StatusInternalServerError)

		return
	}

	response.Success(
		w,
		"User created successfully",
		http.StatusCreated,
		map[string]string{
			"document_id": docRef.ID,
		})
}
