package register

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nova-hrm/app/domain/nova-hrm/middleware"
	"nova-hrm/app/domain/nova-hrm/models"
	"nova-hrm/app/domain/nova-hrm/response"
	"nova-hrm/app/domain/nova-hrm/utils"
	"nova-hrm/app/domain/nova-hrm/validators"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=6,max=50,startswithalpha"`
	Password string `json:"password" validate:"required,min=6,max=50"`
	Email    string `json:"email" validate:"required,email"`
}

type registerResponse struct {
	DocumentId string `json:"document_id"`
}

func HandleAddUser(w http.ResponseWriter, r *http.Request) {
	var registerReq RegisterRequest
	var requestUser *models.User

	requestUser, ok := r.Context().Value(middleware.UserContextKey).(*models.User)
	fmt.Println(requestUser)

	if !ok {
		response.Error(w, "Unauthorized access", http.StatusUnauthorized)

		return
	}

	if requestUser.Role != "admin" {
		response.Error(w, "Unauthorized access", http.StatusUnauthorized)

		return
	}

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

	encPassword, err := utils.HashPassword(registerReq.Password)

	if err != nil {
		response.Error(w, "Error hashing password", http.StatusInternalServerError)

		return
	}

	// Save user to Firestore
	docId, err := createUser(
		registerReq.Username,
		encPassword,
		registerReq.Email,
	)
	if err != nil {
		response.Error(w, "Error creating user", http.StatusInternalServerError)

		return
	}

	data := registerResponse{
		DocumentId: docId,
	}

	response.Success(
		w,
		"User created successfully",
		http.StatusCreated,
		data,
	)
}
