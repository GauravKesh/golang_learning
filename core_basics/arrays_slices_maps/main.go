package main

import "fmt"

func main() {
	// ---------------------------
	// Arrays in Go
	// ---------------------------
	// - Fixed length (size is part of the type)
	// - Elements must all be of the same type
	// - Example: [3]int means an array of exactly 3 integers
	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// ---------------------------
	// Slices in Go
	// ---------------------------
	// - More powerful and flexible than arrays
	// - Dynamic size (can grow/shrink)
	// - Backed by an underlying array
	// - Often used instead of arrays in real programs
	slc := []string{"go", "java", "python"}
	slc = append(slc, "rust") // append adds elements dynamically
	fmt.Println("Slice:", slc)

	// ---------------------------
	// Maps in Go
	// ---------------------------
	// - Key-value pairs (like dictionaries in Python)
	// - Keys must be of a comparable type (string, int, etc.)
	// - Values can be any type
	// - Very efficient for lookups
	mp := map[string]int{"Alice": 25, "Bob": 30}
	mp["Charlie"] = 35 // adding a new key-value pair
	fmt.Println("Map:", mp)
}
