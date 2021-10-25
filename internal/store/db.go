package store

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	*sql.DB
}

func InitStore() (*Db, error) {

	if _, err := os.Stat("./data"); os.IsNotExist(err) {
		if err := os.Mkdir("./data", os.ModePerm); err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite3", "./data/database.db")
	if err != nil {
		return nil, err
	}

	if err := migrate(db); err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func migrate(db *sql.DB) error {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS auto (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			brand TEXT NOT NULL,
			model TEXT NOT NULL,
			price INTEGER CHECK (price > 0) NOT NUll,
			status TEXT CHECK (
				status IN (
			    	'DELIVERED',
					'STORAGED',
					'SALED',
					'REJECT'
				)
			) DEFAULT 'DELIVERED',
			mileage INTEGER CHECK (mileage > 0) NOT NUll
		);
	`)

	return err
}
