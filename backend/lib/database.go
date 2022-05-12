package lib

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectToDataBase(loadedEnv *Enviroment) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", loadedEnv.HOST, loadedEnv.PORT, loadedEnv.USER, loadedEnv.DBNAME, loadedEnv.PASSWORD)
	db, err := gorm.Open(loadedEnv.DIALECT, dbURI)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database")
	}
	DB = db
}
