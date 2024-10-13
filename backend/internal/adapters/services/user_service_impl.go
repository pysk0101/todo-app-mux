package services

import (
	"github.com/pysk0101/todo-app-mux/backend/internal/core/domain"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/ports"
)

type UserServiceImpl struct {
	repo ports.UserRepository
}

func NewUserServiceImpl(repo ports.UserRepository) ports.UserService {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetUser(id string) (*domain.User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) Create(user *domain.User) error {
	err := s.repo.Create(user)
	return err
}

func (s *UserServiceImpl) Update(user *domain.User) error {
	err := s.repo.Update(user)
	return err
}

func (s *UserServiceImpl) Delete(id string) error {
	err := s.repo.Delete(id)
	return err
}
