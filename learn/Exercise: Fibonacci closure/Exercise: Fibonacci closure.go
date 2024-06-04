package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// Initialize the first two Fibonacci numbers
	a, b := 0, 1

	// Return a closure function
	return func() int {
		// Calculate the next Fibonacci number
		next := a
		a, b = b, a+b
		return next
	}
}

func main() {
	// Get the Fibonacci closure
	f := fibonacci()

	// Print the first 10 Fibonacci numbers
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
