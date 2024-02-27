package handlers

import (
	"database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "jurnol/db"
    "jurnol/models"
)

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser models.User
		err := json.NewDecoder(r.Body).Decode(&newUser)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}