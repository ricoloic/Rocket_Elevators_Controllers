package controller

import "time"

// Wait will add a pause/sleep time in the program
func Wait(t int) {
	// time.Sleep(0 * time.Millisecond) // -- uncomment for no wait *the animation won't work
	time.Sleep(time.Duration(t) * time.Second)
}

// Positive gets the absolute value of a number (integer)
func Positive(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

// Remove remove the element at the provided index in the provided list then it return the new list
func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
