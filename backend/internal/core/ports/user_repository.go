package ports

import "github.com/pysk0101/todo-app-mux/backend/internal/core/domain"

type UserRepository interface {
	GetUser(id string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id string) error
}
