package response

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/shaikhjunaidx/url-shortner/database"
	"github.com/shaikhjunaidx/url-shortner/models"
	"github.com/shaikhjunaidx/url-shortner/utils"
)

func SendResponse(c *fiber.Ctx, shortURLID string, requestBody *models.URLRequest) error {
	redisClient := database.CreateClient(1)
	defer redisClient.Close()

	responseBody := createResponseBody(shortURLID, requestBody)
	ipAddress := c.IP()
	rateLimitKey := utils.GenerateRateLimitKey(ipAddress)

	if err := updateRateLimitInfo(redisClient, rateLimitKey, &responseBody); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(responseBody)
}

func createResponseBody(shortURLID string, requestBody *models.URLRequest) models.URLResponse {
	return models.URLResponse{
		URL:                requestBody.URL,
		CustomShort:        os.Getenv("DOMAIN") + "/" + shortURLID,
		Expiry:             requestBody.Expiry,
		RateLimitRemaining: 10, // Initial value, will be updated
		RateLimitReset:     30, // Initial value, will be updated
	}
}

func updateRateLimitInfo(redisClient *redis.Client, rateLimitKey string, responseBody *models.URLResponse) error {
	redisClient.Decr(database.Ctx, rateLimitKey)
	remainingQuota, err := utils.GetRemainingQuota(redisClient, rateLimitKey)
	if err != nil {
		return err
	}
	responseBody.RateLimitRemaining = remainingQuota

	ttl, err := utils.GetTTL(redisClient, rateLimitKey)
	if err != nil {
		return err
	}
	responseBody.RateLimitReset = ttl

	return nil
}
