package databases

import (
	"database/sql"
	"os"
	"time"
)

func GetConnection() *sql.DB {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("mysql", dsn)
	// for repository_test.go we need to hardcode the dsn here rght the one bellow
	// db, err := sql.Open("mysql", "username:password@tcp(localhost:8000)/db_name?parseTime=true&loc=Local")
	// please replace username, password, db_name with your own mysql configuration
	// also make sure your mysql server is running and accessible
	// and i still dunno why the parseTime and loc are needed here, but it just works this way
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
