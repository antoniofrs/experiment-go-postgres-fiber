package routes

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/controller"
	"github.com/gofiber/fiber/v2"
)

func InitBookRoutes(router fiber.Router, controller *controller.BookController) {
	router.Get("/book", controller.FindAll)
	router.Get("/book/:Id", controller.FindById)
	router.Post("/book", controller.Create)
	router.Put("/book/:Id", controller.Update)
	router.Delete("/book/:Id", controller.Delete)
}
