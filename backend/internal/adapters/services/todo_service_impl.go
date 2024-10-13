package services

import (
	"github.com/pysk0101/todo-app-mux/backend/internal/core/domain"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/ports"
)

type TodoServiceImpl struct {
	repo ports.TodoRepository
}

func NewTodoServiceImpl(repo ports.TodoRepository) ports.TodoService {
	return &TodoServiceImpl{repo: repo}
}

func (s *TodoServiceImpl) GetAll() ([]*domain.Todo, error) {
	todos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoServiceImpl) Create(todo *domain.Todo) error {
	err := s.repo.Create(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoServiceImpl) Update(todo *domain.Todo) error {
	err := s.repo.Update(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoServiceImpl) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
