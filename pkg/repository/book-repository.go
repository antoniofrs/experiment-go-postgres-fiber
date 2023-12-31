package repository

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/model"
	"github.com/google/uuid"
)

type BookRepository interface {
	Save(book *model.Book)
	Delete(book *model.Book)
	FindById(id uuid.UUID) *model.Book
	FindAll() *[]model.Book
}
