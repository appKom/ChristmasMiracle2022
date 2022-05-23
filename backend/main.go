package main

import (
	"log"
	"net/http"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/lib"
	"github.com/appKom/ChristmasMiracle2022/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	subRouter := myRouter.PathPrefix("/api/v1").Subrouter()
	authRouter := subRouter.PathPrefix("/auth").Subrouter()

	// For tasks
	subRouter.HandleFunc("/tasks", routes.GetTasks).Methods("GET", "OPTIONS")
	subRouter.HandleFunc("/tasks/{id}", routes.GetTask).Methods("GET", "OPTIONS")

	subRouter.HandleFunc("/tasks", routes.CheckAuthMiddleware(routes.CheckAdminMiddleware(routes.CreateTask))).Methods("POST", "OPTIONS")
	subRouter.HandleFunc("/tasks/{id}", routes.CheckAuthMiddleware(routes.CheckAdminMiddleware(routes.DeleteTask))).Methods("DELETE", "OPTIONS")

	// For debugging purposes, should be removed in production
	subRouter.HandleFunc("/flags", routes.CheckAuthMiddleware(routes.CheckAdminMiddleware(routes.GetFlags))).Methods("GET", "OPTIONS")
	authRouter.HandleFunc("/users", routes.CheckAuthMiddleware(routes.CheckAdminMiddleware(routes.GetUsers))).Methods("GET", "OPTIONS")
	// For submitting
	subRouter.HandleFunc("/submit/{id}", routes.CheckAuthMiddleware(routes.SubmitFlag)).Methods("POST", "OPTIONS")

	subRouter.HandleFunc("/scoreboard", routes.CheckAuthMiddleware(routes.GetScoreBoard)).Methods("GET", "OPTIONS")
	subRouter.HandleFunc("/profile", routes.CheckAuthMiddleware(routes.GetProfile)).Methods("GET", "OPTIONS")

	// For authentication
	authRouter.HandleFunc("/login", routes.LoginUser).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/register", routes.CreateUser).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/refresh", routes.RefreshToken).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/logout", routes.NotImplemented).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	lib.LoadSystemEnv()
	lib.ConnectToDataBase(lib.LoadedEnv)

	defer lib.DB.Close()

	lib.DB.AutoMigrate(&api.Task{})
	lib.DB.AutoMigrate(&api.Flag{})
	lib.DB.AutoMigrate(&api.User{})

	handleRequests()
}
