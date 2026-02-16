package models

import "time"

// Todo represents a todo item
type Todo struct {
	ID          int64     `json:"id" db:"id" example:"1"`
	Title       string    `json:"title" db:"title" validate:"required,min=3,max=100" example:"Buy groceries"`
	Description string    `json:"description,omitempty" db:"description" example:"Milk, eggs, bread"`
	Completed   bool      `json:"completed" db:"completed" example:"false"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" example:"2026-02-16T09:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" example:"2026-02-16T09:00:00Z"`
}

func (Todo) TableName() string {
	return "todos"
}
