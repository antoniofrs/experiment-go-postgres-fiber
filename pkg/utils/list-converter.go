package utils

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/dto"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/model"
)

// TODO: implement a better solution
func ConvertSlice(books []model.Book) []dto.BookDto {

	var booksDto []dto.BookDto

	for _, book := range books {
		bookDto := book.ToDto() 
		booksDto = append(booksDto, bookDto)
	}

	return booksDto
}
