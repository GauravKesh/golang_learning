package main

import "fmt"

// Structs in Go:
// - A struct is a collection of fields (like a class without inheritance).
// - Used to group related data together.
// - Structs are commonly used to model real-world entities (e.g., Person, Car, etc.).

// Defining a struct 'Person' with two fields: Name and Age
type Person struct {
	Name string
	Age  int
}

// Methods in Go:
// - A method is a function with a receiver (the part before the method name).
// - The receiver allows the method to operate on the struct's data.
// - Syntax: func (receiver Type) MethodName(params) returnType
func (p Person) Greet() {
	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.Name, p.Age)
}

func main() {
	// Creating an instance of Person struct
	p := Person{Name: "Gaurav", Age: 21}

	// Printing the struct directly
	fmt.Println("Person Struct:", p)

	// Calling the Greet method on struct 'p'
	p.Greet()
}
