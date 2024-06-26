package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shaikhjunaidx/url-shortner/middleware"
	"github.com/shaikhjunaidx/url-shortner/models"
	"github.com/shaikhjunaidx/url-shortner/response"
	"github.com/shaikhjunaidx/url-shortner/services"
	"github.com/shaikhjunaidx/url-shortner/utils"
)

func ShortenURL(c *fiber.Ctx) error {
	requestBody := new(models.URLRequest)
	if err := utils.ParseRequestBody(c, requestBody); err != nil {
		return err
	}

	if err := middleware.CheckRateLimit(c); err != nil {
		return err
	}

	if err := middleware.ValidateURL(requestBody); err != nil {
		return err
	}

	shortURLID := utils.GenerateShortID(requestBody)

	if err := services.CheckShortIDAvailability(shortURLID); err != nil {
		return err
	}

	utils.SetDefaultExpiry(requestBody)

	if err := services.StoreShortURL(shortURLID, requestBody); err != nil {
		return err
	}

	return response.SendResponse(c, shortURLID, requestBody)
}
