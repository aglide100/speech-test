package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
)


func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, errors.New("failed to ping: " + err.Error())
    }

	return db, nil
}

