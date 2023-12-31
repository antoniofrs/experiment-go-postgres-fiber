package controller

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{BookService: bookService}
}


func (c *BookController) Create(ctx *fiber.Ctx) error {
	panic("Not implemented yet")
}

func (c *BookController) Update(ctx *fiber.Ctx) error {
	panic("Not implemented yet")
}

func (c *BookController) Delete(ctx *fiber.Ctx) error {
	panic("Not implemented yet")
}

func (c *BookController) FindAll(ctx *fiber.Ctx) error {
	panic("Not implemented yet")
}

func (c *BookController) FindById(ctx *fiber.Ctx) error {
	panic("Not implemented yet")
}
