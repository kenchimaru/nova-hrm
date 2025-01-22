package middleware

import (
	// "context"
	"context"
	"net/http"
	"nova-hrm/app/domain/nova-hrm/database"
)

type firestoreContextKey string

var contextKey firestoreContextKey = "firebase"

func FirestoreMiddleware(fc *database.FirestoreClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), contextKey, fc)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
