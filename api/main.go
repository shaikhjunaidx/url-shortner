package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/shaikhjunaidx/url-shortner/handlers"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/:url", handlers.ResolveURL)
	app.Post("/api/v1", handlers.ShortenURL)
}

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setUpRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
