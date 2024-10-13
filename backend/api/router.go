package api

import (
	"github.com/gorilla/mux"
	"github.com/pysk0101/todo-app-mux/backend/internal/adapters/handler"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/ports"
)

func NewRouter(todoService ports.TodoService, userService ports.UserService) *mux.Router {
	// Gorilla Mux router'ı oluştur
	router := mux.NewRouter()

	// Todo handler'ı oluştur
	todoHandler := handler.NewTodoHandler(todoService)
	userHandler := handler.NewUserHandler(userService)

	// Rotaları tanımla
	router.HandleFunc("/todos", todoHandler.GetAll).Methods("GET")
	router.HandleFunc("/todos", todoHandler.Create).Methods("POST")
	router.HandleFunc("/todos/{id}", todoHandler.Update).Methods("PUT")
	router.HandleFunc("/todos/{id}", todoHandler.Delete).Methods("DELETE")

	router.HandleFunc("/users", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	return router
}
