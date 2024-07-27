package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ReadJSON(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteResponse(w, status, map[string]string{"error": message})
}

func WriteErrors(w http.ResponseWriter, status int, errors []error) {
	// Convert errors to a slice of strings
	messages := make([]string, len(errors))
	for i, err := range errors {
		if err != nil {
			messages[i] = err.Error()
		}
	}

	// Create the error response map
	errorResponse := map[string][]string{"errors": messages}

	// Use the WriteJSON function to send the JSON response
	WriteResponse(w, status, errorResponse)
}
