package routes

import "github.com/gofiber/fiber/v2"

func InitHealtRoutes(router fiber.Router) error {

	router.Get("/healt", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	return nil
}
