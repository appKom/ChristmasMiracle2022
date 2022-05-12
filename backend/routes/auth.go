package routes

import (
	"encoding/json"
	"net/http"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/auth"
	"github.com/appKom/ChristmasMiracle2022/lib"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginCredentials api.LoginCredentials
	json.NewDecoder(r.Body).Decode(&loginCredentials)

	var user api.User

	lib.DB.Where(&api.User{Email: loginCredentials.Email}).First(&user)

	if auth.CheckPasswordHash(loginCredentials.Password, user.Password) {
		token, err := auth.CreateNewTokenPair(user, lib.LoadedEnv.JWT_SECRET)
		if err != nil {
			SetHeaders(w, http.StatusInternalServerError)
			return
		}
		SetHeaders(w, http.StatusOK)

		json.NewEncoder(w).Encode(token)
	} else {
		SetHeaders(w, http.StatusBadRequest)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user api.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "" || user.Email == "" || user.Password == "" {
		SetHeaders(w, http.StatusBadRequest)
		return
	}

	hashedPassword, _ := auth.HashPassword(user.Password)
	user.Password = hashedPassword
	user.Points = 0
	user.Admin = false

	created := lib.DB.Create(&user)
	if created.Error != nil {
		SetHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode("An error occured, while creating user")
	}

	createdUser := api.CreatedUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Points:   user.Points,
		Admin:    user.Admin,
	}

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(createdUser)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var refreshToken api.RefreshToken
	json.NewDecoder(r.Body).Decode(&refreshToken)

	sub, err := auth.CheckTokenValidity(refreshToken.Refresh, lib.LoadedEnv.JWT_SECRET)

	if err != nil {
		SetHeaders(w, http.StatusBadRequest)
		return
	}

	var user api.User
	var token api.TokenResponse

	lib.DB.First(&user, sub)

	token, err = auth.CreateNewTokenPair(user, lib.LoadedEnv.JWT_SECRET)
	if err != nil {
		SetHeaders(w, http.StatusInternalServerError)
		return
	}

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
