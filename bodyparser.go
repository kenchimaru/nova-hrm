package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func bodyParserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Unable to parse form", http.StatusBadRequest)

				return
			}

			formData := make(map[string]string)
			for key, values := range r.Form {
				if len(values) > 0 {
					formData[key] = values[0]
				}
			}
			jsonData, err := json.Marshal(formData)
			if err != nil {
				http.Error(w, "Unable to convert form to JSON", http.StatusInternalServerError)

				return
			}

			r.Body = io.NopCloser(bytes.NewBuffer(jsonData))
			r.Header.Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}
