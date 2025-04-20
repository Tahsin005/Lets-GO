package main

import (
	"errors"
	"fmt"
)

func divide (a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func main () {
	result1, err := divide(10.0, 5.0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result1)

	result2, err := divide(10.0, 3.0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result2)

	result3, err := divide(10.0, 0.0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result3)
}
