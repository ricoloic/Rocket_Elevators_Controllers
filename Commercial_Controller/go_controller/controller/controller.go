package controller

import "time"

// Wait will add a pause/sleep time in the program
func Wait(t int) {
	// time.Sleep(0 * time.Millisecond) -- uncomment for no wait *the animation won't work
	time.Sleep(time.Duration(t) * time.Second)
}

// Positive gets the absolute value of a number (integer)
func Positive(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
