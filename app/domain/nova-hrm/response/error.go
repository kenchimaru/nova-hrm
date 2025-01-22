package response

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]any{
		"code":    code,
		"message": msg,
	})
}

func Success(w http.ResponseWriter, msg string, code int, data any) {
	if msg == "" {
		msg = "Success"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]any{
		"code":     code,
		"message":  msg,
		"response": data,
	})
}
