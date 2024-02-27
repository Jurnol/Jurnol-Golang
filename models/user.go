package models

type User struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Email        string `json:"email,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"` // Temporarily included for completeness
	// CreatedAt and UpdatedAt are managed by the db.
}