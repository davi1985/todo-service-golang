package services

import (
	"fmt"
	"strings"

	"todo-api/internal/models"
	"todo-api/internal/repositories"
)

type TodoService interface {
	GetAll() ([]models.Todo, error)
	GetByID(id int64) (*models.Todo, error)
	Create(todo *models.Todo) error
	Update(todo *models.Todo) error
	Delete(id int64) error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) GetAll() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) GetByID(id int64) (*models.Todo, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid id: %d", id)
	}
	
	return s.repo.GetByID(id)
}

func (s *todoService) Create(todo *models.Todo) error {
	if err := s.validateTodo(todo); err != nil {
		return err
	}
	
	todo.Title = strings.TrimSpace(todo.Title)
	todo.Description = strings.TrimSpace(todo.Description)
	
	return s.repo.Create(todo)
}

func (s *todoService) Update(todo *models.Todo) error {
	if err := s.validateTodo(todo); err != nil {
		return err
	}
	
	if todo.ID <= 0 {
		return fmt.Errorf("invalid id for update: %d", todo.ID)
	}
	
	_, err := s.repo.GetByID(todo.ID)
	if err != nil {
		return fmt.Errorf("todo not found for update: %w", err)
	}
	
	todo.Title = strings.TrimSpace(todo.Title)
	todo.Description = strings.TrimSpace(todo.Description)
	
	return s.repo.Update(todo)
}

func (s *todoService) Delete(id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid id for delete: %d", id)
	}
	
	_, err := s.repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("todo not found for delete: %w", err)
	}
	
	return s.repo.Delete(id)
}

func (s *todoService) validateTodo(todo *models.Todo) error {
	if todo == nil {
		return fmt.Errorf("todo cannot be nil")
	}
	
	title := strings.TrimSpace(todo.Title)
	if title == "" {
		return fmt.Errorf("title is required")
	}
	
	if len(title) < 3 {
		return fmt.Errorf("title must be at least 3 characters long")
	}
	
	if len(title) > 100 {
		return fmt.Errorf("title must be less than 100 characters")
	}
	
	if todo.Description != "" {
		description := strings.TrimSpace(todo.Description)
		if len(description) > 500 {
			return fmt.Errorf("description must be less than 500 characters")
		}
	}
	
	return nil
}
