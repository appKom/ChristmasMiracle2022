package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/appKom/ChristmasMiracle2022/lib"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Task struct {
	gorm.Model

	Title   string
	Content string
}

type Flag struct {
	gorm.Model

	Key    string
	TaskID int
}

var db *gorm.DB

func getTasks(w http.ResponseWriter, r *http.Request) {

	var tasks []Task

	db.Find(&tasks)

	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task Task

	db.First(&task, params["id"])

	json.NewEncoder(w).Encode(task)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)

	created := db.Create(&task)
	err := created.Error

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/tasks", getTasks).Methods("GET")
	myRouter.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	myRouter.HandleFunc("/tasks", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {

	loadedEnv := lib.LoadSystemEnv()
	db = lib.ConnectToDataBase(loadedEnv)

	defer db.Close()

	db.AutoMigrate(&Task{}, &Flag{})
	handleRequests()
}
