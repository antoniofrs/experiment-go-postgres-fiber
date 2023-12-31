package controller

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/dto"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (c *BookController) Create(ctx *fiber.Ctx) error {

	insertBookDto := new(dto.InsertBookDto)
	if err := ctx.BodyParser(insertBookDto); err != nil {
		return err
	}

	bookDto := c.BookService.Create(insertBookDto)

	return ctx.Status(fiber.StatusOK).JSON(bookDto)
}

func (c *BookController) Update(ctx *fiber.Ctx) error {

	id, _ := uuid.Parse(ctx.Params("Id"))
	insertBookDto := new(dto.InsertBookDto)
	if err := ctx.BodyParser(insertBookDto); err != nil {
		return err
	}

	bookDto := c.BookService.Update(id, insertBookDto)

	return ctx.Status(fiber.StatusOK).JSON(bookDto)

}

func (c *BookController) Delete(ctx *fiber.Ctx) error {

	id, _ := uuid.Parse(ctx.Params("Id"))
	c.BookService.Delete(id)
	return ctx.SendStatus(fiber.StatusOK)

}

func (c *BookController) FindAll(ctx *fiber.Ctx) error {
	booksDto := c.BookService.FindAll()
	return ctx.Status(fiber.StatusOK).JSON(booksDto)
}

func (c *BookController) FindById(ctx *fiber.Ctx) error {
	id, _ := uuid.Parse(ctx.Params("Id"))

	bookDto := c.BookService.FindById(id)
	return ctx.Status(fiber.StatusOK).JSON(bookDto)

}
