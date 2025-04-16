package main

import "fmt"

func main() {
	// _ ignores the second value that is returned from the function (because go doesn't allow us to use unused variables)
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}

func getNames() (string, string) {
	return "John", "Doe"
}
