package utils 

import (
    "math/rand"
    "time"
)

func Random(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := make([]rune, length)
	for i := range digits {
		digits[i] = rune('0' + rand.Intn(10))
	}
	return string(digits)
 }