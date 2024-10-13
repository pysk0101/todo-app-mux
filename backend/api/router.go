package api

import (
	"github.com/gorilla/mux"
	"github.com/pysk0101/todo-app-mux/backend/internal/adapters/handlers"
	"github.com/pysk0101/todo-app-mux/backend/internal/core/ports"
)

// NewRouter rotaları oluşturur ve middleware'leri ekler.
func NewRouter(router *mux.Router, todoService ports.TodoService, userService ports.UserService, authService ports.AuthService) *mux.Router {
	// Todo handler'ı oluştur
	todoHandler := handlers.NewTodoHandler(todoService)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	// Rotaları tanımla
	router.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// JWT middleware'ini kullanarak korunan rotalar
	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(JWTMiddleware(authService))

	// Todo rotaları
	protected.HandleFunc("/todos", todoHandler.GetAll).Methods("GET")
	protected.HandleFunc("/todos", todoHandler.Create).Methods("POST")
	protected.HandleFunc("/todos/{id}", todoHandler.Update).Methods("PUT")
	protected.HandleFunc("/todos/{id}", todoHandler.Delete).Methods("DELETE")

	// User rotaları
	protected.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	protected.HandleFunc("/users", userHandler.Create).Methods("POST")
	protected.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	protected.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	return router
}
