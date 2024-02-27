package database

import (
	"database/sql"
	_"github.com/lib/pq"
)

func NewConnection() (*sql.DB, error) {
	connStr := "user=antoine dbname=jurnolgolang password=papichulo33 host=localhost port=5432 sslmode=disable"

		db, err := sql.Open("postgres", connStr)

		if err != nil {
			return nil, err
		}

		if err = db.Ping(); err != nil {
			return nil, err
		}

		return db, nil
}