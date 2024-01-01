package main

import (
	"errors"

	"github.com/antoniofrs/experiment-go-postgresql/pkg/controller"
	module "github.com/antoniofrs/experiment-go-postgresql/pkg/module"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/repository"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/routes"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	module.InitConfigs()

	db := module.NewPostgres()

	// Book
	BookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(BookRepository)
	bookController := controller.NewBookController(bookService)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			var e *fiber.Error
			if !errors.As(err, &e) {
				e.Code = fiber.StatusInternalServerError
				e.Message = "Internal Server Error"
			}

			ctx.Status(e.Code).SendString(e.Message)

			return nil
		},
	})

	app.Use(recover.New())

	api := app.Group("/api")

	routes.InitBookRoutes(api, bookController)
	routes.InitHealtRoutes(api)

	app.Listen(":3000")

}
