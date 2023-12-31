package model

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/dto"
	"github.com/google/uuid"
)

type Book struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	Title            string
	Author           string
	PubblicationDate string
	Completed        bool
}

func (book Book) ToDto() dto.BookDto {
	return dto.BookDto{
		ID:               uuid.Nil,
		Title:            book.Title,
		Author:           book.Author,
		PubblicationDate: book.PubblicationDate,
		Completed:        book.Completed,
	}
}