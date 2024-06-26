package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shaikhjunaidx/url-shortner/models"
)

func ParseRequestBody(c *fiber.Ctx, requestBody *models.URLRequest) error {
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	return nil
}
