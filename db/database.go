package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect(driverName, dbName string) error {
	db, err := sql.Open(driverName, dbName)
	if err != nil {
		return fmt.Errorf("db.Connect() error #1: \t\n %w \r\n", err)
	}

	DB = db
	return nil
}
