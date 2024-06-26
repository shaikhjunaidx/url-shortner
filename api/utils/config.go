package utils

import (
	"os"
)

func GetAPIQuota() string {
	return os.Getenv("API_QUOTA")
}
