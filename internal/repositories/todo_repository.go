package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"todo-api/internal/database"
	"todo-api/internal/models"
)

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	GetByID(id int64) (*models.Todo, error)
	Create(todo *models.Todo) error
	Update(todo *models.Todo) error
	Delete(id int64) error
}

type todoRepository struct {
	db *database.DB
}

func NewTodoRepository(db *database.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) GetAll() ([]models.Todo, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at 
		FROM todos 
		ORDER BY created_at DESC
	`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query todos: %w", err)
	}
	defer rows.Close()
	
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		var description sql.NullString
		
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&description,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		
		if description.Valid {
			todo.Description = description.String
		}
		
		todos = append(todos, todo)
	}
	
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	
	return todos, nil
}

func (r *todoRepository) GetByID(id int64) (*models.Todo, error) {
	query := `
		SELECT id, title, description, completed, created_at, updated_at 
		FROM todos 
		WHERE id = ?
	`
	
	var todo models.Todo
	var description sql.NullString
	
	err := r.db.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Title,
		&description,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("todo with id %d not found", id)
		}
		return nil, fmt.Errorf("failed to query todo by id: %w", err)
	}
	
	if description.Valid {
		todo.Description = description.String
	}
	
	return &todo, nil
}

func (r *todoRepository) Create(todo *models.Todo) error {
	query := `
		INSERT INTO todos (title, description, completed) 
		VALUES (?, ?, ?)
	`
	
	var description interface{}
	if todo.Description != "" {
		description = todo.Description
	}
	
	result, err := r.db.Exec(query, todo.Title, description, todo.Completed)
	if err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}
	
	todo.ID = id
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	
	return nil
}

func (r *todoRepository) Update(todo *models.Todo) error {
	query := `
		UPDATE todos 
		SET title = ?, description = ?, completed = ? 
		WHERE id = ?
	`
	
	var description interface{}
	if todo.Description != "" {
		description = todo.Description
	}
	
	result, err := r.db.Exec(query, todo.Title, description, todo.Completed, todo.ID)
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %d not found", todo.ID)
	}
	
	todo.UpdatedAt = time.Now()
	
	return nil
}

func (r *todoRepository) Delete(id int64) error {
	query := `DELETE FROM todos WHERE id = ?`
	
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("todo with id %d not found", id)
	}
	
	return nil
}
