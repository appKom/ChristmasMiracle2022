package routes

import "net/http"

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, http.StatusNotImplemented)
}
