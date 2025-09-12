package main

import (
	"fmt"
)

func learnArithmeticOperator() {
	//  addition
	const a int64 = 56
	const b int64 = 54

	const add int64 = a + b
	var minus int64 = a - b
	var mult int64 = a * b

	fmt.Println(add)
	fmt.Println(minus)
	fmt.Println(mult)
}

func learnControlFlow() {
	const num int = 3

	variable := "string"

	for i := 0; i < num; i++ {
		fmt.Print(i)
		fmt.Printf("the user name is %v", variable)
	}

}

func main() {
	learnArithmeticOperator()
	learnControlFlow()
}
