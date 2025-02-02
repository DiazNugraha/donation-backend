package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoute(app *fiber.App) {
	route := app.Group("/swagger")

	route.Get("*", swagger.HandlerDefault)
}
