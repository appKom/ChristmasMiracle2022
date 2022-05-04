package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/lib"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func getTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []api.Task

	db.Find(&tasks)

	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task api.Task

	db.First(&task, params["id"])

	json.NewEncoder(w).Encode(task)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task api.Task
	json.NewDecoder(r.Body).Decode(&task)

	created := db.Create(&task)
	err := created.Error

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task api.Task

	db.First(&task, params["id"])
	db.Delete(&task)

	json.NewEncoder(w).Encode(task)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	subRouter := myRouter.PathPrefix("/api/v1").Subrouter()
	subRouter.HandleFunc("/tasks", getTasks).Methods("GET")
	subRouter.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	subRouter.HandleFunc("/tasks", createTask).Methods("POST")
	subRouter.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	loadedEnv := lib.LoadSystemEnv()
	db = lib.ConnectToDataBase(loadedEnv)

	defer db.Close()

	db.AutoMigrate(&api.Task{}, &api.Flag{})
	handleRequests()
}
