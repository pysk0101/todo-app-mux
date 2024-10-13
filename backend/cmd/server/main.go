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
	"github.com/pysk0101/todo-app-mux/backend/internal/utils"
)

func main() {
	// .env dosyasını yükle
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file: %s\n", err)
		os.Exit(1)
	}

	// Veritabanını başlat
	db.InitDB()
	utils.RunMigrations()

	// Repository ve Service'leri oluştur
	todoRepository := repositories.NewTodoRepositoryImpl(db.GetDB())
	todoService := services.NewTodoServiceImpl(todoRepository)

	userRepository := repositories.NewUserRepositoryImpl(db.GetDB())
	userService := services.NewUserServiceImpl(userRepository)

	authService := services.NewAuthServiceImpl(userRepository) // AuthService'i oluştur

	// Router'ı oluştur
	router := mux.NewRouter()

	// Router'ı yapılandır
	router = api.NewRouter(router, todoService, userService, authService)

	// Port ayarlarını al
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Sunucuyu başlat
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %s\n", err)
		os.Exit(1)
	}
}
