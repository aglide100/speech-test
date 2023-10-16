package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*Database, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, errors.New("failed to ping: " + err.Error())
    }

	return &Database{
		conn: db,
	}, nil
}

type Database struct {
	conn *sql.DB
}