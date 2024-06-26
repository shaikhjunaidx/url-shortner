package models

import "time"

type URLRequest struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type URLResponse struct {
	URL                string        `json:"url"`
	CustomShort        string        `json:"short"`
	Expiry             time.Duration `json:"expiry"`
	RateLimitRemaining int           `json:"rate_limit"`
	RateLimitReset     time.Duration `json:"rate_limit_reset"`
}
