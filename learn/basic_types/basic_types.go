package main

import (
	"fmt"
	"math/cmplx"
)

// Declare package-level variables
var (
	ToBe   bool       = false                // Boolean variable
	MaxInt uint64     = 1<<64 - 1            // Unsigned integer variable
	z      complex128 = cmplx.Sqrt(-5 + 12i) // Complex number variable
)

func main() {
	// Print the type and value of each variable
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)     // %T prints the type of the variable, %v prints its value
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) // %T prints the type of the variable, %v prints its value
	fmt.Printf("Type: %T Value: %v\n", z, z)            // %T prints the type of the variable, %v prints its value
}
