package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/auth"
	"github.com/appKom/ChristmasMiracle2022/lib"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

func checkAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("No token provided")
			return
		}

		extractedToken := strings.Split(token, "Bearer ")
		if len(extractedToken) == 2 {
			token = strings.TrimSpace(extractedToken[1])
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Incorrect token format")
			return
		}

		email, err := auth.CheckTokenValidity(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Invalid token")
			return
		}

		r.Header.Set("Email", email)
		next(w, r)
	})
}

// Get all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []api.Task

	db.Find(&tasks)

	json.NewEncoder(w).Encode(tasks)
}

// Get task by ID
func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task api.Task

	db.First(&task, params["id"])

	json.NewEncoder(w).Encode(task)
}

// Create new task with key
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask api.NewTask
	json.NewDecoder(r.Body).Decode(&newTask)

	if newTask.Title == "" || newTask.Content == "" || newTask.Points == 0 || newTask.Key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var task api.Task
	var flag api.Flag

	task.Title = newTask.Title
	task.Content = newTask.Content
	task.Points = newTask.Points

	tx := db.Begin()

	created := tx.Create(&task)
	err := created.Error
	if err != nil {
		tx.Rollback()
		json.NewEncoder(w).Encode(err)
	}

	flag.Key = newTask.Key // Should be hashed
	flag.TaskID = task.ID

	created = tx.Create(&flag)
	err = created.Error
	if err != nil {
		tx.Rollback()
		json.NewEncoder(w).Encode(err)
	}

	tx.Commit()

	json.NewEncoder(w).Encode(task)
}

// Delete task, with flag
func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task api.Task
	var flag api.Flag

	db.First(&task, params["id"])
	db.Where(&api.Flag{TaskID: task.ID}).First(&flag)

	tx := db.Begin()

	err := tx.Exec("DELETE FROM tasks WHERE id = ?", task.ID).Error
	if err != nil {
		tx.Rollback()
		json.NewEncoder(w).Encode(err)
	}

	err = tx.Exec("DELETE FROM flags WHERE id = ?", flag.ID).Error
	if err != nil {
		tx.Rollback()
		json.NewEncoder(w).Encode(err)
	}
	tx.Commit()

	json.NewEncoder(w).Encode(task)
}

// Submit flag and check if its correct
func submitFlag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var submittedFlag api.SubmittedFlag
	var flag api.Flag
	var task api.Task

	db.First(&task, params["id"])
	json.NewDecoder(r.Body).Decode(&submittedFlag)

	db.Where(&api.Flag{Key: submittedFlag.Key}).First(&flag)

	if flag.Key == submittedFlag.Key {
		// TODO: award points to user
		json.NewEncoder(w).Encode(flag)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// Get all flags - debug
func getFlags(w http.ResponseWriter, r *http.Request) {
	var flags []api.Flag

	db.Find(&flags)

	json.NewEncoder(w).Encode(flags)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user api.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "" || user.Email == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, _ := hashPassword(user.Password)
	user.Password = hashedPassword
	user.Points = 0

	created := db.Create(&user)
	err := created.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	var createdUser api.CreatedUser
	createdUser.ID = user.ID
	createdUser.Username = user.Username
	createdUser.Email = user.Email
	createdUser.Points = user.Points

	json.NewEncoder(w).Encode(createdUser)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var loginCredentials api.LoginCredentials
	json.NewDecoder(r.Body).Decode(&loginCredentials)

	var user api.User

	db.Where(&api.User{Email: loginCredentials.Email}).First(&user)

	if checkPasswordHash(loginCredentials.Password, user.Password) {
		token, err := auth.CreateNewToken(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(token)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []api.User

	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	subRouter := myRouter.PathPrefix("/api/v1").Subrouter()
	authRouter := subRouter.PathPrefix("/auth").Subrouter()

	// For tasks
	subRouter.HandleFunc("/tasks", getTasks).Methods("GET")
	subRouter.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	subRouter.HandleFunc("/tasks", createTask).Methods("POST")
	subRouter.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	// For debugging purposes, should be removed in production
	subRouter.HandleFunc("/flags", getFlags).Methods("GET")
	authRouter.HandleFunc("/users", checkAuthMiddleware(getUsers)).Methods("GET")

	// For submitting
	subRouter.HandleFunc("/submit/{id}", submitFlag).Methods("POST")

	// For authentication
	authRouter.HandleFunc("/login", loginUser).Methods("POST")
	authRouter.HandleFunc("/register", createUser).Methods("POST")
	authRouter.HandleFunc("/logout", notImplemented).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	loadedEnv := lib.LoadSystemEnv()
	db = lib.ConnectToDataBase(loadedEnv)

	defer db.Close()

	db.AutoMigrate(&api.Task{}, &api.Flag{}, &api.User{})
	handleRequests()
}
