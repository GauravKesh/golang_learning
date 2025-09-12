package main

import "fmt"

// Functions in Go:
// - A function is a block of code that performs a specific task.
// - It can take parameters and return values.
// - Functions help in code reuse, modularity, and clarity.

// Simple function: takes two integers and returns their sum.
func add(a int, b int) int {
	return a + b
}

// Multiple return values in Go:
// - Functions can return multiple values.
// - Commonly used for returning both result and error status.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// Returning an error if denominator is zero.
		return 0, fmt.Errorf("division by zero")
	}
	// Returning quotient and no error (nil).
	return a / b, nil
}

func main() {
	// Calling add function
	fmt.Println("Sum:", add(5, 7))

	// Calling divide function with error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}
}
