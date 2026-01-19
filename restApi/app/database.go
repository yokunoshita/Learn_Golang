package app

import (
	"database/sql"
	"restApi/helper"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=yoku password=yokunoshita dbname=restapi_golang sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}
