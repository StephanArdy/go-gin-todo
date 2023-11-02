package models

import "time"

type Todo struct {
	ID        int     `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	Title     string    `json:"title" form:"title" binding:"required"`
	Completed bool      `json:"completed"`
	DueDate   string `json:"due_date" form:"due_date" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
}
