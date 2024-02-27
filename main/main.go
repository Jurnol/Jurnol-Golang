package main

import (

	"database/sql"
	_"github.com/lib/pq"
	// "jurnol/db"
	"log"
)

func main() {
	connStr := "user=antoine dbname=jurnolgolang password=papichulo33 host=localhost port=5432 sslmode=disable"

	db, err:= sql.Open("postgres", connStr)
	if err !=nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var currentTime string
    query := "SELECT NOW()"
    err = db.QueryRow(query).Scan(&currentTime)
    if err != nil {
        log.Fatal("Failed to execute query: ", err)
    }

	log.Printf("Successfully connected to database. Current time: %v", currentTime)
	

}