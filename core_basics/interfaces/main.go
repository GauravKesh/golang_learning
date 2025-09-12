package main

import "fmt"

// ---------------------------
// Interfaces in Go
// ---------------------------
// - An interface defines a set of method signatures.
// - Any type that provides implementations for those
//   methods is said to "satisfy" the interface.
// - No explicit declaration is needed (no "implements"
//   keyword like in Java). It's implicit.
// - This enables polymorphism and abstraction.
// ---------------------------

// Define an interface with one method
type Greeter interface {
	Greet()
}

// Struct that will implement the interface
type Person struct {
	Name string
}

// Method that matches the interface definition
// Since Person has a Greet() method, it automatically
// implements the Greeter interface.
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	// Interface variable can hold any type
	// that satisfies the interface.
	var g Greeter = Person{Name: "Gaurav"}

	// Calls the implementation of Greet() defined on Person
	g.Greet()
}
