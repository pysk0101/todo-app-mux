package ports

import "github.com/pysk0101/todo-app-mux/backend/internal/core/domain"

type AuthService interface {
	Register(user *domain.User) error
	Login(username, password string) (string, error) // JWT döndürür
}
