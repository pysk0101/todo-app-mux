package api

import (
	"github.com/gorilla/mux"
	"github.com/pysk0101/todo-app-mux/backend/internal/adapters/handler"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/ports"
)

func NewRouter(todoService ports.TodoService) *mux.Router {
	// Gorilla Mux router'ı oluştur
	router := mux.NewRouter()

	// Todo handler'ı oluştur
	todoHandler := handler.NewTodoHandler(todoService)

	// Rotaları tanımla
	router.HandleFunc("/todos", todoHandler.GetAll).Methods("GET")
	router.HandleFunc("/todos", todoHandler.Create).Methods("POST")
	router.HandleFunc("/todos/{id}", todoHandler.Update).Methods("PUT")
	router.HandleFunc("/todos/{id}", todoHandler.Delete).Methods("DELETE")

	return router
}
