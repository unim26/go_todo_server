package netservice

import (
	"encoding/json"
	"net/http"
)

//send response
func SendResponse(w http.ResponseWriter, statusCode int, message string, data any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"status": statusCode,
		"message": message,
		"data": data,
	})
}
