package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shaikhjunaidx/url-shortner/database"
	"github.com/shaikhjunaidx/url-shortner/models"
)

func CheckShortIDAvailability(shortURLID string) error {
	redisClient := database.CreateClient(0)
	defer redisClient.Close()

	existingURL, _ := redisClient.Get(database.Ctx, shortURLID).Result()
	if existingURL != "" {
		return fiber.NewError(fiber.StatusForbidden, "Custom Shortened URL is already in use")
	}
	return nil
}

func StoreShortURL(shortURLID string, requestBody *models.URLRequest) error {
	redisClient := database.CreateClient(0)
	defer redisClient.Close()

	err := redisClient.Set(database.Ctx, shortURLID, requestBody.URL, requestBody.Expiry).Err()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to connect to server")
	}
	return nil
}
