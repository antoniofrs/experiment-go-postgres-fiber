package routes

import "github.com/gofiber/fiber/v2"

func HealtController(router fiber.Router) error {

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	return nil
}
