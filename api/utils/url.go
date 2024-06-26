package utils

import (
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shaikhjunaidx/url-shortner/models"
)

func GenerateShortID(requestBody *models.URLRequest) string {
	if requestBody.CustomShort == "" {
		return uuid.New().String()[:6]
	}
	return requestBody.CustomShort
}

func SetDefaultExpiry(requestBody *models.URLRequest) {
	if requestBody.Expiry == 0 {
		requestBody.Expiry = 24 * time.Hour
	}
}

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	return newURL != os.Getenv("DOMAIN")
}
