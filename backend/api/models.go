package api

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model

	Title   string
	Content string
	Points  int
}

type Flag struct {
	gorm.Model

	Key    string
	TaskID uint
}

type User struct {
	gorm.Model

	Username string
	Email    string `gorm:"unique"`
	Password string
	Points   int
	Admin    bool
}
