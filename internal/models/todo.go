package models

import "time"

type Todo struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title" validate:"required,min=3,max=100"`
	Description string    `json:"description,omitempty" db:"description"`
	Completed   bool      `json:"completed" db:"completed"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func (Todo) TableName() string {
	return "todos"
}
