package services

import (
	"github.com/shaikhjunaidx/url-shortner/database"
)

func GetOriginalURL(shortURL string) (string, error) {
	redisClient := database.CreateClient(0)
	defer redisClient.Close()

	return redisClient.Get(database.Ctx, shortURL).Result()
}

func IncrementCounter() error {
	redisClient := database.CreateClient(1)
	defer redisClient.Close()

	return redisClient.Incr(database.Ctx, "counter").Err()
}
