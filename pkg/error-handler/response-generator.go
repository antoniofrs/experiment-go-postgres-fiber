package errorhandler

import "github.com/gofiber/fiber/v2"

func InternalServerException(message string) *fiber.Error {
	return &fiber.Error{Code: fiber.StatusInternalServerError, Message: message}
}

func NotFoundException(message string) *fiber.Error {
	return &fiber.Error{Code: fiber.StatusNotFound, Message: message}
}
