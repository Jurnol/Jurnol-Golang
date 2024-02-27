package main

import (

	"database/sql"
	_"github.com/lib/pq"
	// "jurnol/db"
	"log"
)

func main() {
	//Connection string to db
	connStr := "user=antoine dbname=jurnolgolang password=papichulo33 host=localhost port=5432 sslmode=disable"

	//Open fucntion comes from the database/sql package - needs a first arg of db driver - postgres - and connection string arg
	db, err:= sql.Open("postgres", connStr)
	if err !=nil {
		log.Fatal(err)
	}

	defer db.Close()
	//Pings the db to make sure connection is alive. Will return an error if it ain't
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var currentTime string
    query := "SELECT NOW()"
		//QueryRow is for an query that is expected to return only one row
		//Per our tutorial experience, the .Scan method is used for to store the value of current time in memore using a pointer
    err = db.QueryRow(query).Scan(&currentTime)
    if err != nil {
        log.Fatal("Failed to execute query: ", err)
    }

	log.Printf("Successfully connected to database. Current time: %v", currentTime)
	

}