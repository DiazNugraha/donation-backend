package main

import (
	"donation-backend/pkg/configs"
	"donation-backend/pkg/middlewares"
	"donation-backend/pkg/routes"
	"donation-backend/pkg/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := configs.FiberConfig()

	app := fiber.New(config)

	middlewares.FiberMiddleware(app)

	routes.SwaggerRoute(app)
	routes.NotFoundRoute(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}

}
