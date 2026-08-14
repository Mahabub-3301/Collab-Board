package main

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}

	return nil
}

func writeError(w http.ResponseWriter, status int, message string) error {
	return writeJSON(w, status, APIError{
		Error: message,
	})
}n
