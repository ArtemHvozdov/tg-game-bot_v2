package models

// Struct for storing task information loaded from JSON
type Task struct {
	ID          int    `json:"id"`
	Title      string `json:"title"` // NB: "Tittle" likely meant to be "Title"
	Description string `json:"description"`
}
