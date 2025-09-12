package main

import "fmt"

func main() {
	// Variables in Go:
	// - Declared with 'var' or shorthand ':='.
	// - Statically typed, type can be explicit or inferred.
	var x int = 42
	y := 100 // shorthand declaration (Go infers type int)

	// Constants:
	// - Declared with 'const'.
	// - Values cannot be changed after declaration.
	const pi = 3.1415

	// Pointers in Go:
	// - Store memory addresses of values.
	// - '&' gives the address of a variable.
	// - '*' dereferences a pointer (gets the value stored at the address).
	ptr := &x
	fmt.Println("x:", x)                 // prints value of x
	fmt.Println("Pointer address:", ptr) // prints memory address of x
	fmt.Println("Pointer value:", *ptr)  // dereferences pointer to get value of x

	// Printing other variables
	fmt.Println("y:", y, "pi:", pi)
}
