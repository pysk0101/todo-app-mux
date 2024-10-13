package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pysk0101/todo-app-mux/backend/internal/core/domain"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/ports"
)

type TodoHandler struct {
	service ports.TodoService
}

type ITodoHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewTodoHandler(service ports.TodoService) ITodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var todo domain.Todo

	json.NewDecoder(r.Body).Decode(&todo)

	err := h.service.Create(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	var todo domain.Todo

	json.NewDecoder(r.Body).Decode(&todo)

	err := h.service.Update(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
