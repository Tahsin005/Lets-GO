package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// curry function
func enhanceFunction(functionToModify func(int, int) int) func (int) int {
	return func (x int) int {
		return functionToModify(x, x)
	} 
}

func main () {
	// modifying function using currying
	// currying is a technique in functional programming where a function with multiple arguments is transformed into a sequence of functions, each taking a single argument.
	squareFunction := enhanceFunction(multiply)
	doubleFunction := enhanceFunction(add)

	fmt.Println("Square:", squareFunction(5)) // Output: 25
	fmt.Println("Double:", doubleFunction(5)) // Output: 10
}