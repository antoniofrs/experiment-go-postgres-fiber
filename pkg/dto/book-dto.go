package dto

import (
	"github.com/google/uuid"
)

type BookDto struct {
	ID               uuid.UUID `json:"id"`
	Title            string    `json:"title"`
	Author           string    `json:"author"`
	PubblicationDate string    `json:"pubblicationDate"`
	Completed        bool      `json:"completed"`
}

type InsertBookDto struct {
	Title            string `json:"title"`
	Author           string `json:"author"`
	PubblicationDate string `json:"pubblicationDate"`
}