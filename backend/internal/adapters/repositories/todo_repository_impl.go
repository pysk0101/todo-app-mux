package repositories

import (
	"database/sql"
	"errors"

	"github.com/pysk0101/todo-app-mux/backend/internal/core/domain"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/ports"
)

type TodoRepositoryImpl struct {
	db *sql.DB
}

func NewTodoRepositoryImpl(db *sql.DB) ports.TodoRepository {
	return &TodoRepositoryImpl{db: db}
}

func (r *TodoRepositoryImpl) GetAll() ([]*domain.Todo, error) {
	rows, err := r.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := []*domain.Todo{}
	for rows.Next() {
		t := &domain.Todo{}
		err := rows.Scan(&t.Id, &t.Title, &t.IsDone, &t.Created_At)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (r *TodoRepositoryImpl) Create(todo *domain.Todo) error {
	_, err := r.db.Exec("INSERT INTO todos (title) VALUES (?)", todo.Title)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepositoryImpl) Update(todo *domain.Todo) error {
	_, err := r.db.Exec("UPDATE todos SET title = ?, is_done = ? WHERE id = ?", todo.Title, todo.IsDone, todo.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TodoRepositoryImpl) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
