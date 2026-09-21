package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func sendJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func SendSuccess(w http.ResponseWriter, status int, message string, data any) {
	sendJSON(w, status, Response{
		Code:    status,
		Message: message,
		Data:    data,
	})
}

func SendError(w http.ResponseWriter, status int, message string) {
	sendJSON(w, status, Response{
		Code:    status,
		Message: message,
	})
}
