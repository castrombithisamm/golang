package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    // Initial guess for square root
    z := 1.0
    
    // Iterate until a good approximation is found
    for i := 0; i < 10; i++ {
        // Update the guess based on Newton's method
        z -= (z*z - x) / (2*z)
        // Print the guess at each iteration
        fmt.Printf("Iteration %d: z = %v\n", i+1, z)
    }
    
    return z
}

func main() {
    // Test the square root function for various values of x
    for x := 1.0; x <= 10.0; x++ {
        fmt.Printf("Square root of %v:\n", x)
        fmt.Println("Our function:", Sqrt(x))
        fmt.Println("Standard library function:", math.Sqrt(x))
        fmt.Println()
    }
}
