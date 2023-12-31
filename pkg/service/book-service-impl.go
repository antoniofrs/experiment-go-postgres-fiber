package service

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/dto"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/model"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/repository"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/utils"
	"github.com/google/uuid"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookService(BookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{BookRepository: BookRepository}
}

// Create implements BookService.
func (s *BookServiceImpl) Create(InsertBookDto dto.InsertBookDto) dto.BookDto {

	book := model.Book{
		ID:               uuid.Nil,
		Title:            InsertBookDto.Title,
		Author:           InsertBookDto.Author,
		PubblicationDate: InsertBookDto.PubblicationDate,
	}

	s.BookRepository.Save(&book)

	return book.ToDto()

}

// FindAll implements BookService.
func (s *BookServiceImpl) FindAll() []dto.BookDto {
	books := s.BookRepository.FindAll()
	return utils.ConvertSlice(*books)
}

// FindById implements BookService.
func (s *BookServiceImpl) FindById(id uuid.UUID) dto.BookDto {

	book := s.BookRepository.FindById(id)
	return book.ToDto()

}

// Update implements BookService.
func (s *BookServiceImpl) Update(id uuid.UUID, InsertBookDto dto.InsertBookDto) dto.BookDto {

	book := s.BookRepository.FindById(id)

	if book == nil {
		panic("No books with id " + id.String())
	}

	newBook := model.Book{
		ID:               book.ID,
		Title:            InsertBookDto.Title,
		Author:           InsertBookDto.Author,
		PubblicationDate: InsertBookDto.PubblicationDate,
	}

	s.BookRepository.Save(&newBook)

	return newBook.ToDto()
}

// Delete implements BookService.
func (s *BookServiceImpl) Delete(id uuid.UUID) {

	book := s.BookRepository.FindById(id)

	if book == nil {
		panic("No books with id " + id.String())
	}

	s.BookRepository.Delete(book)

}
