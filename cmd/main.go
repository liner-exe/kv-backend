package main

import "net/http"

func GetUserMe(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "ok", http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/user/me", GetUserMe)
}
