package database

import (
    "testing"
    "database/sql"
    _ "github.com/lib/pq"
)

func TestDatabaseConnection(t *testing.T) {
    connStr := "user=antoine dbname=jurnolgolang password=papichulo33 host=localhost port=5432 sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        t.Fatalf("Failed to open a DB connection: %v", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        t.Fatalf("Failed to connect to the database: %v", err)
    }
}