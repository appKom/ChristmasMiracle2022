package routes

import "net/http"

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, http.StatusNotImplemented)
}

// function that sets headers
func SetHeaders(w http.ResponseWriter, status int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(status)
}
