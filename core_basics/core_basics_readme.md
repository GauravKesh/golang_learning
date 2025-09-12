# Go Core Language Basics 🚀

This repository contains standalone Go programs that demonstrate the **core concepts of the Go programming language**.  
Each concept is placed in its own folder with a `main.go` file, making it easy to run and learn step by step.

---

## 📂 Folder Structure
```bash
├── core_basics
│   ├── arrays_slices_maps
│   │   └── main.go
│   ├── control_structures
│   │   └── main.go
│   ├── core_basics_readme.md
│   ├── functions
│   │   └── main.go
│   ├── interfaces
│   │   └── main.go
│   ├── structs_methods
│   │   └── main.go
│   ├── syntax
│   │   └── main.go
│   └── variables_constants_pointers
│       └── main.go
```

---

## 🚀 How to Run

Navigate into any example folder and run the program:

```bash
cd core-basics/01_syntax
go run main.go
```
Repeat for other folders (variables_constants_pointers, functions, etc.).

# 📘 Core Concepts

## 1. Syntax & Data Types
Go has a **statically typed, simple syntax**. Common data types include:
- `int`, `float64` → numeric types
- `string` → text
- `bool` → boolean values (true/false)

Each program starts with a `main` package and `main()` function.

---

## 2. Variables, Constants & Pointers
- **Variables**: Declared with `var` or shorthand `:=`.  
- **Constants**: Declared with `const` and cannot be changed.  
- **Pointers**: Hold memory addresses (`&` for address, `*` for dereferencing`).  

Pointers help avoid unnecessary copies and allow direct memory manipulation.

---

## 3. Functions & Multiple Return Values
- Functions are declared with the `func` keyword.  
- Go supports **multiple return values**, useful for returning results and errors.  
- Error handling is explicit in Go (`if err != nil { ... }`).  

This keeps code predictable and clear.

---

## 4. Control Structures
Go provides:
- `if` and `else` (no parentheses required)  
- `for` loop (the only loop in Go, can replace `while`)  
- `switch` for cleaner branching logic  

These make flow control concise and readable.

---

## 5. Arrays, Slices & Maps
- **Arrays**: Fixed-length sequences (`[3]int{1,2,3}`).  
- **Slices**: Dynamic, more common than arrays (`[]string{"a","b"}`).  
- **Maps**: Key-value pairs (`map[string]int{"Alice":25}`).  

Slices and maps are central to most Go programs.

---

## 6. Structs & Methods
- **Structs**: Custom data types that group fields together.  
- **Methods**: Functions with a receiver, bound to a struct type.  

This is Go’s way of achieving OOP-like behavior (without classes).

---

## 7. Interfaces
- Define **behavior** by specifying method signatures.  
- Any type that implements those methods automatically satisfies the interface.  

Interfaces allow **polymorphism** and are key for abstractions in Go.
