package routes

import (
	"encoding/json"
	"net/http"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/lib"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []api.User

	lib.DB.Find(&users)
	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("sub")
	var user api.User
	var solvedTasks []api.Task

	lib.DB.First(&user, id)

	lib.DB.Model(&user).Related(&solvedTasks, "SolvedTasks")

	var userResponse api.CreatedUser

	userResponse.ID = user.ID
	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Points = user.Points
	userResponse.Admin = user.Admin

	for _, task := range solvedTasks {
		userResponse.SolvedTasks = append(userResponse.SolvedTasks, task.ID)
	}

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}

func GetScoreBoard(w http.ResponseWriter, r *http.Request) {
	type ScoreBoard struct {
		Username string `json:"Username"`
		Points   int    `json:"Points"`
	}

	var users []ScoreBoard
	lib.DB.Raw("SELECT username, points FROM users ORDER BY points desc").Scan(&users)

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
