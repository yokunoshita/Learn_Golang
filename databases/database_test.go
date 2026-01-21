package databases

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func TestOpenConnect(t *testing.T) {

	dsn := os.Getenv("DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal("Failed to open database connection:", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal("Failed to ping database:", err)
	}
	defer db.Close()
}
