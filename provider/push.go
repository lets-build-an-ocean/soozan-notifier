package provider

import (
	"math/rand"
	"time"
)

func PushSMS(to string, text string) int {
	// Sleep a short time to simulate network delay
	time.Sleep(200 * time.Millisecond)

	// 20% chance of failure
	if rand.Float64() < 0.2 {
		return 500
	}

	return 200
}
