package lib

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Enviroment struct {
	DIALECT    string
	USER       string
	PASSWORD   string
	DBNAME     string
	HOST       string
	PORT       string
	JWT_SECRET string
}

var LoadedEnv *Enviroment

func LoadSystemEnv() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	host := os.Getenv("DB_HOST")
	dialect := os.Getenv("DB_DIALECT")
	port := os.Getenv("DB_PORT")
	jwtSecret := os.Getenv("JWT_SECRET")

	LoadedEnv = &Enviroment{
		DIALECT:    dialect,
		USER:       user,
		PASSWORD:   password,
		DBNAME:     dbName,
		HOST:       host,
		PORT:       port,
		JWT_SECRET: jwtSecret,
	}
}
