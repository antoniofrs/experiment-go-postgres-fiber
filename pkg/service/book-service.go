package service

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/dto"
	"github.com/google/uuid"
)

type BookService interface {
	Create(InsertBookDto dto.InsertBookDto) dto.BookDto
	Update(Id uuid.UUID, InsertBookDto dto.InsertBookDto) dto.BookDto
	Delete(Id uuid.UUID)
	FindById(Id uuid.UUID) dto.BookDto
	FindAll() []dto.BookDto
}
