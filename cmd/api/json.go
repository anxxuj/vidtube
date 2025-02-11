package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]any

// Sends a JSON response with the given status code and headers.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Add additional headers to response if provided.
	for key, val := range headers {
		w.Header()[key] = val
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(append(js, '\n'))

	return nil
}
