package main

import (
	// no longer using this because it w
	// "database/sql"
	"net/http"
	_"github.com/lib/pq"
	"jurnol/db"
	"jurnol/handlers"
	"log"
)

func main() {
	//Connection string to db
	connStr := "user=antoine dbname=jurnolgolang password=papichulo33 host=localhost port=5432 sslmode=disable"

	//Open fucntion comes from the database/sql package - needs a first arg of db driver - postgres - and connection string arg
	database, err:= database.NewConnection(connStr)
	if err !=nil {
		log.Fatal(err)
	}


	http.HandleFunc("/register", handlers.CreateUserHandler(database))

	log.Println("Server starting on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	// defer database.Close()
	// //Pings the db to make sure connection is alive. Will return an error if it ain't
	// err = database.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var currentTime string
  //   query := "SELECT NOW()"
	// 	//QueryRow is for an query that is expected to return only one row
	// 	//Per our tutorial experience, the .Scan method is used for to store the value of current time in memore using a pointer
  //   err = database.QueryRow(query).Scan(&currentTime)
  //   if err != nil {
  //       log.Fatal("Failed to execute query: ", err)
  //   }

	// log.Printf("Successfully connected to database. Current time: %v", currentTime)
	

}