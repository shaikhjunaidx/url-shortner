package middleware

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/shaikhjunaidx/url-shortner/models"
	"github.com/shaikhjunaidx/url-shortner/utils"
)

func ValidateURL(requestBody *models.URLRequest) error {
	if !govalidator.IsURL(requestBody.URL) {
		return fiber.NewError(fiber.StatusBadRequest, "Not a valid URL")
	}
	if !utils.RemoveDomainError(requestBody.URL) {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Domain error")
	}
	requestBody.URL = utils.EnforceHTTP(requestBody.URL)
	return nil
}
