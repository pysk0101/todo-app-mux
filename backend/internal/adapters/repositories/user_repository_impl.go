package repositories

import (
	"database/sql"
	"errors"

	"github.com/pysk0101/todo-app-mux/backend/internal/core/domain"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/ports"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) ports.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetUser(id string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, email, password) VALUES (?)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Update(user *domain.User) error {
	_, err := r.db.Exec("UPDATE users SET username = ?, email = ? , password= ? WHERE id = ?", user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
