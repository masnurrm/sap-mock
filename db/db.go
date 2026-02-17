package db

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/lib/pq"
)

func NewConnection() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "sapuser"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "sappassword"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "sapdb"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
	return sql.Open("postgres", dsn)
}