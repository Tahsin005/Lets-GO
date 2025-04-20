package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func add(x, y int, name string) int {
	// if the same type of arguments are in order, the we could declare the type on the last one if needed
	fmt.Println(name)
	return x + y
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")

	fmt.Println(add(5, 7, "tahsin"))
}

