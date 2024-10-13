package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var Database *sql.DB

func InitDB() {
	var dbUrl = os.Getenv("DATABASE_URL")
	var dbToken = os.Getenv("DATABASE_TOKEN")

	if dbUrl == "" || dbToken == "" {
		fmt.Fprintln(os.Stderr, "DATABASE_URL or DATABASE_TOKEN is not set")
		os.Exit(1)
	}

	url := fmt.Sprintf("%v?authToken=%v", dbUrl, dbToken)

	var err error
	Database, err = sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open db: %s\n", err)
		os.Exit(1)
	}

	if err = Database.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to the database: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to database successfully")

}

func GetDB() *sql.DB {
	return Database
}
