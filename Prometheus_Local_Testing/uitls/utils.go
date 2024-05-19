package myuitls

import (
	"math/rand"
	"time"
)

func RandomInt(min, max int) int {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())
	// Generate a random integer in the range [min, max)
	return min + rand.Intn(max-min)
}
