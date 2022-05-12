package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/lib"
	"github.com/gorilla/mux"
)

// Get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []api.Task

	lib.DB.Where("release_date < ?", time.Now()).Find(&tasks)

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// Get task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task api.Task

	lib.DB.Where("release_date < ?", time.Now()).First(&task, params["id"])

	if task.ID == 0 {
		SetHeaders(w, http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// Create new task with key
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask api.NewTask
	json.NewDecoder(r.Body).Decode(&newTask)

	if newTask.Title == "" || newTask.Content == "" || newTask.Points == 0 || newTask.Key == "" || newTask.ReleaseDate == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	date, date_err := time.Parse(time.RFC3339, newTask.ReleaseDate)
	if date_err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task := api.Task{
		Title:       newTask.Title,
		Content:     newTask.Content,
		Points:      newTask.Points,
		ReleaseDate: date,
	}
	tx := lib.DB.Begin()

	created := tx.Create(&task)
	err := created.Error
	if err != nil {
		tx.Rollback()
		SetHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	flag := api.Flag{
		Key:    newTask.Key,
		TaskID: task.ID,
	}

	created = tx.Create(&flag)
	err = created.Error
	if err != nil {
		tx.Rollback()
		SetHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	tx.Commit()

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// Delete task, with flag
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task api.Task
	var flag api.Flag

	lib.DB.First(&task, params["id"])
	lib.DB.Where(&api.Flag{TaskID: task.ID}).First(&flag)

	tx := lib.DB.Begin()

	err := tx.Exec("DELETE FROM tasks WHERE id = ?", task.ID).Error
	if err != nil {
		tx.Rollback()
		SetHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	err = tx.Exec("DELETE FROM flags WHERE id = ?", flag.ID).Error
	if err != nil {
		tx.Rollback()
		SetHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
	tx.Commit()

	SetHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
