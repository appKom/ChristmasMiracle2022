package routes

import (
	"encoding/json"
	"net/http"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/lib"
)

// Gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []api.User

	lib.DB.Find(&users)
	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// Gets user by ID
func GetProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("sub")

	if id == "" {
		SetHeaders(w, http.StatusBadRequest)
		json.NewEncoder(w).Encode("No user ID provided")
		return
	}

	var user api.User
	var solvedTasks []api.Task

	lib.DB.First(&user, id)
	lib.DB.Model(&user).Related(&solvedTasks, "SolvedTasks")

	userResponse := api.CreatedUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Points:   user.Points,
		Admin:    user.Admin,
	}

	for _, task := range solvedTasks {
		userResponse.SolvedTasks = append(userResponse.SolvedTasks, task.ID)
	}

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}

// Gets scoreboard
func GetScoreBoard(w http.ResponseWriter, r *http.Request) {
	var users []api.ScoreBoard
	lib.DB.Raw("SELECT username, points FROM users ORDER BY points desc").Scan(&users)

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
