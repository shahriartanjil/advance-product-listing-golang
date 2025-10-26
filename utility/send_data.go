package utility

import (
	"encoding/json"
	"net/http"
)

// SendData sends JSON response with given status code and data
func SendData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, `{"error":"failed to encode response"}`, http.StatusInternalServerError)
	}
}

// SendError sends JSON error message with given status code
func SendError(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := map[string]string{
		"error": msg,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, `{"error":"failed to encode error response"}`, http.StatusInternalServerError)
	}
}
