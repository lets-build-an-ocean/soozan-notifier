package provider

import (
	"gosms/logger"
	"math/rand"
	"time"
)

func PushSMS(to string, text string) int {
	logger.Debug("Initiating SMS push request to: " + to)

	// Sleep a short time to simulate network delay
	time.Sleep(200 * time.Millisecond)

	// 20% chance of failure
	if rand.Float64() < 0.2 {
		logger.Warn("SMS provider returned error for: " + to)
		return 500
	}

	logger.Debug("SMS sent successfully to: " + to)
	return 200
}
