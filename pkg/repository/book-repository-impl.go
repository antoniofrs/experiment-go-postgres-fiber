package repository

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	Db *gorm.DB
}

func NewBookRepository(Db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

func (b *BookRepositoryImpl) Delete(book *model.Book) {

	result := b.Db.Delete(book)

	if (result.Error) != nil {
		panic("dbError")
	}
}

// FindAll implements BookRepository.
func (b *BookRepositoryImpl) FindAll() *[]model.Book {
	
	var books []model.Book
	result := b.Db.Find(&books)

	if (result.Error) != nil {
		panic("dbError")
	}

	return &books
	
}

// FindById implements BookRepository.
func (b *BookRepositoryImpl) FindById(bookId uuid.UUID) *model.Book {
	book := new(model.Book)
	
	if b.Db.First(book, bookId).Error != nil { 
		panic("dbError")
	}

	if(book.ID == uuid.Nil){
		return nil
	}
	
	return book
}

// Save implements BookRepository.
func (b *BookRepositoryImpl) Save(book *model.Book) {

	result := b.Db.Save(book)

	if (result.Error) != nil {
		panic(result.Error)
	}
}
