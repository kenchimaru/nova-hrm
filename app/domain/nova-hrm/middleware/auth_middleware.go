package middleware

import (
	"context"
	"fmt"
	"net/http"
	"nova-hrm/app/config"
	"nova-hrm/app/domain/nova-hrm/response"

	"github.com/golang-jwt/jwt/v5"
)

// Define a custom type for context keys
type contextKey string

const UserContextKey contextKey = "user"

// Middleware to verify JWT
func Auth(next http.Handler) http.Handler {
	var jwtSecret = []byte(config.GetEnv("JWT_SECRET", "secret"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			response.Error(w, "Missing token", http.StatusUnauthorized)

			return
		}

		tokenString = tokenString[7:] // Remove "Bearer " prefix

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is correct
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			response.Error(w, "Invalid token", http.StatusUnauthorized)

			return
		}

		userID := token.Claims.(jwt.MapClaims)["user_id"].(string)

		if userID == "" {
			response.Error(w, "Invalid token", http.StatusUnauthorized)

			return
		}

		user, err := getUserByID(userID)
		if err != nil {
			response.Error(w, "Error fetching user", http.StatusInternalServerError)

			return
		}

		if user == nil {
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
