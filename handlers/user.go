package handlers

import (
	"database/sql"
    "encoding/json"
    // "fmt"
    "log"
    "net/http"
    // "jurnol/db"
    "jurnol/models"
)

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser models.UserRequest
		err := json.NewDecoder(r.Body).Decode(&newUser)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		query := `INSERT INTO users (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4)`
		_, err = db.Exec(query, newUser.User.FirstName, newUser.User.LastName, newUser.User.Email, newUser.User.PasswordHash) 
		if err != nil {
			log.Printf("Failed to insert new user: %v", err)
			http.Error(w, "failed to create new user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message" : " User created successfully"})
	}
}