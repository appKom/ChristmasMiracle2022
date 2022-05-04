package api

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model

	Title   string
	Content string
	Points int
}

type Flag struct {
	gorm.Model

	Key    string 
	TaskID uint 
}

