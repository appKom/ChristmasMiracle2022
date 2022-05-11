package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/auth"
	"github.com/appKom/ChristmasMiracle2022/lib"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/exp/slices"
)

var db *gorm.DB
var jwtSecret string

// Auth middleware
func checkAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if r.Method == "OPTIONS" {
			setHeaders(w, http.StatusOK)
			return
		}

		if token == "" {
			setHeaders(w, http.StatusUnauthorized)
			json.NewEncoder(w).Encode("No token provided")
			return
		}

		extractedToken := strings.Split(token, "Bearer ")
		if len(extractedToken) == 2 {
			token = strings.TrimSpace(extractedToken[1])
		} else {
			setHeaders(w, http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Incorrect token format")
			return
		}

		UID, err := auth.CheckTokenValidityWithClaims(token, jwtSecret)
		if err != nil {
			setHeaders(w, http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Invalid token")
			return
		}

		r.Header.Set("sub", fmt.Sprint(UID))
		next(w, r)
	})
}

// Admin middleware
func checkAdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UID := r.Header.Get("sub")

		var user api.User
		db.First(&user, UID)

		if !user.Admin {
			setHeaders(w, http.StatusUnauthorized)
			json.NewEncoder(w).Encode("You are not admin")
			return
		} else {
			next(w, r)
		}
	})
}

// function that sets headers
func setHeaders(w http.ResponseWriter, status int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(status)
}

// Get all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []api.Task

	db.Where("release_date < ?", time.Now()).Find(&tasks)

	setHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// Get task by ID
func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task api.Task

	db.Where("release_date < ?", time.Now()).First(&task, params["id"])

	if task.ID == 0 {
		setHeaders(w, http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	setHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// Create new task with key
func createTask(w http.ResponseWriter, r *http.Request) {
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
	tx := db.Begin()

	created := tx.Create(&task)
	err := created.Error
	if err != nil {
		tx.Rollback()
		setHeaders(w, http.StatusInternalServerError)
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
		setHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	tx.Commit()

	setHeaders(w, http.StatusOK)
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
		setHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	err = tx.Exec("DELETE FROM flags WHERE id = ?", flag.ID).Error
	if err != nil {
		tx.Rollback()
		setHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
	tx.Commit()

	setHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// Submit flag and check if its correct
func submitFlag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var submittedFlag api.SubmittedFlag
	var flag api.Flag
	var user api.User
	var task api.Task

	db.First(&user, r.Header.Get("sub"))
	json.NewDecoder(r.Body).Decode(&submittedFlag)

	db.First(&task, params["id"])

	if task.ID == 0 {
		setHeaders(w, http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid task ID")
		return
	}

	db.Where(&api.Flag{Key: submittedFlag.Key, TaskID: task.ID}).First(&flag)

	if flag.ID == 0 {
		setHeaders(w, http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid flag")
		return
	}
	// TODO: award points to user

	var solvedTasks []api.Task
	db.Model(&user).Related(&solvedTasks, "SolvedTasks")

	if slices.Contains(solvedTasks, task) {
		setHeaders(w, http.StatusBadRequest)
		json.NewEncoder(w).Encode("You already solved this task")
		return
	}

	tx := db.Begin()

	err := tx.Model(&user).Association("SolvedTasks").Append(&task).Error
	if err != nil {
		tx.Rollback()
		setHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = tx.Model(&user).Update("Points", user.Points+task.Points).Error

	if err != nil {
		tx.Rollback()
		setHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	tx.Commit()

	setHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(flag)
}

// Get all flags - debug
func getFlags(w http.ResponseWriter, r *http.Request) {
	var flags []api.Flag

	db.Find(&flags)

	setHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(flags)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user api.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "" || user.Email == "" || user.Password == "" {
		setHeaders(w, http.StatusBadRequest)
		return
	}

	hashedPassword, _ := auth.HashPassword(user.Password)
	user.Password = hashedPassword
	user.Points = 0
	user.Admin = false

	created := db.Create(&user)
	err := created.Error
	if err != nil {
		setHeaders(w, http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	var createdUser api.CreatedUser
	createdUser.ID = user.ID
	createdUser.Username = user.Username
	createdUser.Email = user.Email
	createdUser.Points = user.Points
	createdUser.Admin = user.Admin

	setHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(createdUser)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var loginCredentials api.LoginCredentials
	json.NewDecoder(r.Body).Decode(&loginCredentials)

	var user api.User

	db.Where(&api.User{Email: loginCredentials.Email}).First(&user)

	if auth.CheckPasswordHash(loginCredentials.Password, user.Password) {
		token, err := auth.CreateNewTokenPair(user, jwtSecret)
		if err != nil {
			setHeaders(w, http.StatusInternalServerError)
			return
		}
		setHeaders(w, http.StatusOK)

		json.NewEncoder(w).Encode(token)
	} else {
		setHeaders(w, http.StatusBadRequest)
	}
}

func refreshToken(w http.ResponseWriter, r *http.Request) {
	var refreshToken api.RefreshToken
	json.NewDecoder(r.Body).Decode(&refreshToken)

	sub, err := auth.CheckTokenValidity(refreshToken.Refresh, jwtSecret)

	if err != nil {
		setHeaders(w, http.StatusBadRequest)
		return
	}

	var user api.User
	var token api.TokenResponse

	db.First(&user, sub)

	token, err = auth.CreateNewTokenPair(user, jwtSecret)
	if err != nil {
		setHeaders(w, http.StatusInternalServerError)
		return
	}

	setHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []api.User

	db.Find(&users)
	setHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("sub")
	var user api.User
	var solvedTasks []api.Task

	db.First(&user, id)

	db.Model(&user).Related(&solvedTasks, "SolvedTasks")

	var userResponse api.CreatedUser

	userResponse.ID = user.ID
	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Points = user.Points
	userResponse.Admin = user.Admin

	for _, task := range solvedTasks {
		userResponse.SolvedTasks = append(userResponse.SolvedTasks, task.ID)
	}

	setHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}

func getScoreBoard(w http.ResponseWriter, r *http.Request) {
	type ScoreBoard struct {
		Username string `json:"Username"`
		Points   int    `json:"Points"`
	}

	var users []ScoreBoard
	db.Raw("SELECT username, points FROM users ORDER BY points desc").Scan(&users)

	setHeaders(w, http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	setHeaders(w, http.StatusNotImplemented)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	subRouter := myRouter.PathPrefix("/api/v1").Subrouter()
	authRouter := subRouter.PathPrefix("/auth").Subrouter()

	// For tasks
	subRouter.HandleFunc("/tasks", getTasks).Methods("GET", "OPTIONS")
	subRouter.HandleFunc("/tasks/{id}", getTask).Methods("GET", "OPTIONS")

	subRouter.HandleFunc("/tasks", checkAuthMiddleware(checkAdminMiddleware(createTask))).Methods("POST", "OPTIONS")
	subRouter.HandleFunc("/tasks/{id}", checkAuthMiddleware(checkAdminMiddleware(deleteTask))).Methods("DELETE", "OPTIONS")

	// For debugging purposes, should be removed in production
	subRouter.HandleFunc("/flags", checkAuthMiddleware(checkAdminMiddleware(getFlags))).Methods("GET", "OPTIONS")
	authRouter.HandleFunc("/users", checkAuthMiddleware(checkAdminMiddleware(getUsers))).Methods("GET", "OPTIONS")
	// For submitting
	subRouter.HandleFunc("/submit/{id}", checkAuthMiddleware(submitFlag)).Methods("POST", "OPTIONS")

	subRouter.HandleFunc("/scoreboard", checkAuthMiddleware(getScoreBoard)).Methods("GET", "OPTIONS")
	subRouter.HandleFunc("/profile", checkAuthMiddleware(getUser)).Methods("GET", "OPTIONS")

	// For authentication
	authRouter.HandleFunc("/login", loginUser).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/register", createUser).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/refresh", refreshToken).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/logout", notImplemented).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	loadedEnv := lib.LoadSystemEnv()
	db = lib.ConnectToDataBase(loadedEnv)
	jwtSecret = loadedEnv.JWT_SECRET

	defer db.Close()

	db.AutoMigrate(&api.Task{})
	db.AutoMigrate(&api.Flag{})
	db.AutoMigrate(&api.User{})

	handleRequests()
}
