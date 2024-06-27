package routes

import (
	"encoding/json"
	"net/http"
)

func writeError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"error": err.Error(),
	}

	json.NewEncoder(w).Encode(data)
}

func writeResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}

func writeResponseWithStatusCode(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}
