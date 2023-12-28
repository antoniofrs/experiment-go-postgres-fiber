package module

import (
	"github.com/antoniofrs/experiment-go-postgresql/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes() {
	app := fiber.New()
	api := app.Group("/api")

	healt := api.Group("/healt")
	routes.HealtController(healt)

	app.Listen(":3000")
}