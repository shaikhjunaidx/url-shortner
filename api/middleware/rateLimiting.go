package middleware

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/shaikhjunaidx/url-shortner/database"
	"github.com/shaikhjunaidx/url-shortner/utils"
)

func CheckRateLimit(c *fiber.Ctx) error {
	redisClient := database.CreateClient(1)
	defer redisClient.Close()

	ipAddress := c.IP()
	apiQuota := utils.GetAPIQuota()
	rateLimitKey := utils.GenerateRateLimitKey(ipAddress)

	remainingQuota, err := utils.GetRemainingQuota(redisClient, rateLimitKey)

	if err != nil {
		if err == redis.Nil {
			return setInitialQuota(redisClient, rateLimitKey, apiQuota)
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get remaining quota")
	}

	return enforceRateLimit(c, redisClient, rateLimitKey, remainingQuota)
}

func setInitialQuota(redisClient *redis.Client, rateLimitKey, apiQuota string) error {
	err := redisClient.Set(database.Ctx, rateLimitKey, apiQuota, 30*time.Minute).Err()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to set initial quota")
	}
	return nil
}

func enforceRateLimit(c *fiber.Ctx, redisClient *redis.Client, rateLimitKey string, remainingQuota int) error {
	if remainingQuota <= 0 {
		ttl, _ := utils.GetTTL(redisClient, rateLimitKey)
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Rate limit exceeded", "rate_limit_reset": ttl})
	}
	return nil
}
