package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitStore() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./data/database.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
