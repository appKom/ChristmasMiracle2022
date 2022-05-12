package routes

import (
	"encoding/json"
	"net/http"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/lib"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"
)

// Submit flag and check if its correct
func SubmitFlag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var submittedFlag api.SubmittedFlag
	var flag api.Flag
	var user api.User
	var task api.Task

	lib.DB.First(&user, r.Header.Get("sub"))
	json.NewDecoder(r.Body).Decode(&submittedFlag)

	lib.DB.First(&task, params["id"])

	if task.ID == 0 {
		SetHeaders(w, http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid task ID")
		return
	}

	lib.DB.Where(&api.Flag{Key: submittedFlag.Key, TaskID: task.ID}).First(&flag)

	if flag.ID == 0 {
		SetHeaders(w, http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid flag")
		return
	}
	// TODO: award points to user

	var solvedTasks []api.Task
	lib.DB.Model(&user).Related(&solvedTasks, "SolvedTasks")

	if slices.Contains(solvedTasks, task) {
		SetHeaders(w, http.StatusBadRequest)
		json.NewEncoder(w).Encode("You already solved this task")
		return
	}

	tx := lib.DB.Begin()

	err := tx.Model(&user).Association("SolvedTasks").Append(&task).Error
	if err != nil {
		tx.Rollback()
		SetHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = tx.Model(&user).Update("Points", user.Points+task.Points).Error

	if err != nil {
		tx.Rollback()
		SetHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	tx.Commit()

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(flag)
}

// Get all flags - debug
func GetFlags(w http.ResponseWriter, r *http.Request) {
	var flags []api.Flag

	lib.DB.Find(&flags)

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(flags)
}
