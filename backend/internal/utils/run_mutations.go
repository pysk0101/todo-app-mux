package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/pysk0101/todo-app-mux/backend/internal/adapters/db"
)

func RunMigrations() {
	// Migration dosyasını oku
	migrationFile, err := os.ReadFile("../../../backend/internal/adapters/db/migrations/create_tables.sql")
	if err != nil {
		log.Fatalf("Error reading migration file: %v", err)
	}

	// Veritabanında migration'ı çalıştır
	if _, err := db.Database.Exec(string(migrationFile)); err != nil {
		log.Fatalf("Error executing migration: %v", err)
	}

	fmt.Println("Migrations completed successfully")
}
