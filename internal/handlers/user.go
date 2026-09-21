package handlers

import (
	"kv-backend/internal/helpers"
	"net/http"
)

func GetUserMe(w http.ResponseWriter, r *http.Request) {
	user := map[string]any{
		"id":       1,
		"username": "alex",
		"email":    "alex@example.com",
	}

	helpers.SendSuccess(w, http.StatusOK, "success", user)
}
