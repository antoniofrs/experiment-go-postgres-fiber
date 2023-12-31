package model

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/dto"
	"github.com/google/uuid"
)

type Book struct {
	ID               uuid.UUID `gorm:"primaryKey;gorm:type:uuid;default:gen_random_uuid()"`
	Title            string
	Author           string
	PubblicationDate string
}

func (Book) TableName() string {
    return "book"
}

func (book Book) ToDto() dto.BookDto {
	return dto.BookDto{
		ID:               book.ID,
		Title:            book.Title,
		Author:           book.Author,
		PubblicationDate: book.PubblicationDate,
	}
}