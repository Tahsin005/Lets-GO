package main

import "fmt"

type Printer interface {
	Print()
}

type Greeter interface {
	Greet()
}

type Person struct {
	name string
}

func (p Person) Print() {
	fmt.Println("Name:", p.name)
}

func (p Person) Greet() {
	fmt.Println("Hello,", p.name)
}

func main() {
	person := Person{name: "Tahsin"}

	person.Print()   // Output: Name: Tahsin
	person.Greet()   // Output: Hello, Tahsin
}
