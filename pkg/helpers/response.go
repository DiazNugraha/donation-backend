package helpers

import "github.com/gofiber/fiber/v2"

func ResponseJson(status int, message string, data interface{}) fiber.Map {
	return fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	}
}
