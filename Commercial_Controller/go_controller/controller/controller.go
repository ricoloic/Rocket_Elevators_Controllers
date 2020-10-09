package controller

import "time"

// Wait ...
func Wait(t time.Duration) {
	// time.Sleep(t * time.Second)
	time.Sleep(t * time.Millisecond)
}

// Positive ..
func Positive(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
