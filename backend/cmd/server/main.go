package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pysk0101/todo-app-mux/backend/api"
	"github.com/pysk0101/todo-app-mux/backend/internal/adapters/db"
	"github.com/pysk0101/todo-app-mux/backend/internal/adapters/repositories"
	"github.com/pysk0101/todo-app-mux/backend/internal/adapters/services"
)

func main() {
	// .env dosyasını yükle
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file: %s\n", err)
		os.Exit(1)
	}

	db.InitDB()

	todoRepository := repositories.NewTodoRepositoryImpl(db.GetDB())
	todoService := services.NewTodoServiceImpl(todoRepository)

	router := mux.NewRouter()

	api.NewRouter(todoService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %s\n", err)
		os.Exit(1)
	}
}
