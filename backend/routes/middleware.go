package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/auth"
	"github.com/appKom/ChristmasMiracle2022/lib"
)

func CheckAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if r.Method == "OPTIONS" {
			SetHeaders(w, http.StatusOK)
			return
		}

		if token == "" {
			SetHeaders(w, http.StatusUnauthorized)
			json.NewEncoder(w).Encode("No token provided")
			return
		}

		extractedToken := strings.Split(token, "Bearer ")
		if len(extractedToken) == 2 {
			token = strings.TrimSpace(extractedToken[1])
		} else {
			SetHeaders(w, http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Incorrect token format")
			return
		}

		UID, err := auth.CheckTokenValidityWithClaims(token, lib.LoadedEnv.JWT_SECRET)
		if err != nil {
			SetHeaders(w, http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Invalid token")
			return
		}

		r.Header.Set("sub", fmt.Sprint(UID))
		next(w, r)
	})
}

// Admin middleware
func CheckAdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UID := r.Header.Get("sub")

		var user api.User
		lib.DB.First(&user, UID)

		if !user.Admin {
			SetHeaders(w, http.StatusUnauthorized)
			json.NewEncoder(w).Encode("You are not admin")
			return
		} else {
			next(w, r)
		}
	})
}

// function that sets headers
func SetHeaders(w http.ResponseWriter, status int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(status)
}
