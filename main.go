package main

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/controller"
	module "github.com/antoniofrs/experiment-go-postgresql/pkg/module"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/repository"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/routes"
	"github.com/antoniofrs/experiment-go-postgresql/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	module.InitConfigs()
	
	db := module.NewPostgres()

	// Book
	BookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(BookRepository)
	bookController := controller.NewBookController(bookService)


	app := fiber.New()
	api := app.Group("/api")

	routes.InitBookRoutes(api, bookController)
	routes.InitHealtRoutes(api)

	app.Listen(":3000")

}
