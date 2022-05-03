package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Task struct {
	gorm.Model

	Title string
	Content string 
}

type Flag struct {
	gorm.Model

	Key string
	TaskID int
}


func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	//Load env from docker	
	dialect := os.Getenv("POSTGRES_DIALECT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	//Connect to database
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	fmt.Println(dbURI)
	db, err := gorm.Open(dialect, dbURI)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	
	//handleRequests()

	defer db.Close()

	db.AutoMigrate(&Task{}, &Flag{})
}