package main

import "fmt"

// While-style loop in Go
// Go doesn’t have a dedicated "while" keyword.
// Instead, a for-loop with only a condition works as while.
func whileStyle() {
	x := 1
	for x <= 5 { // runs until condition is false
		fmt.Println("While-style loop, x =", x)
		x++
	}
}

func main() {
	x := 10
	y := 20

	// -------------------------------
	// If-Else:
	// - No parentheses required around conditions.
	// - Curly braces are mandatory.
	if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is greater or equal to y")
	}

	// -------------------------------
	// For Loop:
	// - The only looping construct in Go (no while/do-while).
	// - Traditional for with init, condition, and post.
	for i := 1; i <= 3; i++ {
		fmt.Println("For-loop iteration:", i)
	}

	// -------------------------------
	// Switch Statement:
	// - Cleaner alternative to multiple if-else blocks.
	// - Cases don’t require “break”; Go automatically breaks after each case.
	switch x {
	case 5:
		fmt.Println("x is 5")
	case 10:
		fmt.Println("x is 10")
	default:
		fmt.Println("x is something else")
	}

	// -------------------------------
	// While-style loop
	whileStyle()
}
