package main

import (
	"fmt"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/appKom/ChristmasMiracle2022/auth"
	"github.com/appKom/ChristmasMiracle2022/lib"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	lib.LoadSystemEnv()
	lib.ConnectToDataBase(lib.LoadedEnv)

	defer lib.DB.Close()

	lib.DB.AutoMigrate(&api.User{})

	var user api.User

	user.Admin = true

	fmt.Println("Username: ")
	fmt.Scanln(&user.Username)

	fmt.Println("Email: ")
	fmt.Scanln(&user.Email)

	fmt.Println("Password: ")
	fmt.Scanln(&user.Password)

	var confirm string
	fmt.Println("Confirm password: ")
	fmt.Scanln(&confirm)

	if user.Password != confirm {
		fmt.Println("Passwords do not match")
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	user.Password = hashedPassword

	created := lib.DB.Create(&user)
	if created.Error != nil {
		fmt.Println(created.Error)
		return
	}

	fmt.Printf("User created, with username %s\n", user.Username)
}
