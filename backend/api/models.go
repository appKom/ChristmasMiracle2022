package api

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model

	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Points      int       `json:"points"`
	ReleaseDate time.Time `json:"release_date"`
}

type Flag struct {
	gorm.Model

	Key    string `json:"key"`
	TaskID uint   `json:"task_id"`
}

type User struct {
	gorm.Model

	Username    string `json:"username"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	Points      int    `json:"points"`
	Admin       bool   `json:"admin"`
	SolvedTasks []Task `json:"solved_tasks" gorm:"many2many:solved_tasks"`
}
