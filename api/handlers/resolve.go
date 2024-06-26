package handlers

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/shaikhjunaidx/url-shortner/services"
)

func ResolveURL(c *fiber.Ctx) error {
	shortURL := c.Params("url")

	originalURL, err := services.GetOriginalURL(shortURL)
	if err != nil {
		if err == redis.Nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Shortened URL not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot connect to Database"})
	}

	if err := services.IncrementCounter(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to increment counter"})
	}

	return c.Redirect(originalURL, 301)
}
