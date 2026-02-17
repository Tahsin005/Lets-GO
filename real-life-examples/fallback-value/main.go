package main

import (
	"cmp"
	"fmt"
	"slices"
)


func getPortFromEnv() string {
	return "42069"
}

type Employee struct {
	name string
	age  int
}

func main() {
	// example: 1
	port := cmp.Or(
		getPortFromEnv(),
		"8080",
	)
	fmt.Println("Server starting on port: ", port)

	// example: 2
	employees := []Employee{
		{"tahsin", 24},
		{"rean", 22},
		{"ferdous", 24},
	}

	fmt.Println("Employees list before sorting")
	fmt.Println(employees)

	fmt.Println("Employees after sorting")
	fmt.Println(sortEmployees(employees))
}

func sortEmployees(employees []Employee) []Employee {
	sortedEmployees := employees

	slices.SortFunc(sortedEmployees, func(a, b Employee) int {
		return  cmp.Or(
			cmp.Compare(a.age, b.age),
			cmp.Compare(a.name, b.name),
		)
	})

	return sortedEmployees
}

/*
	Compare returns

	-1 if x is less then y
	 0 if x equals y,
	+1 if x is greater than y
*/