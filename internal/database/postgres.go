package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Driver de Postgres
)

// NewPostgresConnection opens a connection to the database
func NewPostgresConnection(host, port, user, password, dbname string) (*sql.DB, error) {
	// Connection string (DSN)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Connection Pool settings (Optimization for Data Specialists)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify the connection is alive
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
