package api

import "github.com/jinzhu/gorm"

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
