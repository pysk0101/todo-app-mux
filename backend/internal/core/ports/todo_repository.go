package ports

import "github.com/pysk0101/todo-app-mux/backend/internal/core/domain"

type TodoRepository interface {
	GetAll() ([]*domain.Todo, error)
	Create(todo *domain.Todo) error
	Update(todo *domain.Todo) error
	Delete(id string) error
}
