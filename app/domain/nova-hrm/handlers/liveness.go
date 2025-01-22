package liveness

import (
	"net/http"
)

// Liveness handles the liveness probe
func Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
