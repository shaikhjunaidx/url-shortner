package utils

import (
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/shaikhjunaidx/url-shortner/database"
)

func GenerateRateLimitKey(ipAddress string) string {
	return ipAddress + ":rate_limit"
}

func GetRemainingQuota(redisClient *redis.Client, rateLimitKey string) (int, error) {
	remainingQuotaStr, err := redisClient.Get(database.Ctx, rateLimitKey).Result()
	if err != nil {
		return 0, err
	}
	remainingQuota, err := strconv.Atoi(remainingQuotaStr)
	if err != nil {
		return 0, err
	}
	return remainingQuota, nil
}

func GetTTL(redisClient *redis.Client, rateLimitKey string) (time.Duration, error) {
	ttl, err := redisClient.TTL(database.Ctx, rateLimitKey).Result()
	if err != nil {
		return 0, err
	}
	return ttl / time.Minute, nil
}
