package main

import "fmt"

func main() {
	// Basic Data Types in Go:
	// Go is statically typed, meaning each variable has a specific type.

	// Integer (int): whole numbers, size depends on system (32-bit or 64-bit).
	var a int = 10

	// Float (float64): decimal numbers, default type is float64.
	var b float64 = 3.14

	// String: sequence of characters, immutable in Go.
	var c string = "Hello Go"

	// Boolean (bool): represents true or false values.
	var d bool = true

	// Printing values using fmt.Println
	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("String:", c)
	fmt.Println("Boolean:", d)
}
