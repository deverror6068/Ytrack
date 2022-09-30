package main

import (
	"fmt"
	"math"
)

// Main function
func Sqrt(nb int) int {

	// Finding square root
	// of the given number
	// Using Sqrt() function
	nvalue_1 := math.Sqrt(3)

	res := nvalue_1
	fmt.Printf("%.3f = %.3f",
		nvalue_1, res)
	return true
}
