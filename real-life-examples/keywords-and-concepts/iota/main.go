package main

import "fmt"

// the default value of iota is 0

const (
	Monday = iota + 1
	Tuesday 
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func main() {
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	fmt.Printf("%03b\n", Readable)
	fmt.Printf("%03b\n", Writeable)
	fmt.Printf("%03b\n", Executable)
}