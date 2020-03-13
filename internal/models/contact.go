package models

// Contact represent a contact object
type Contact struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Email     []string `json:"email"`
}
