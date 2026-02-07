package main

import "fmt"

func main() {
	// pointer
	i, j :=33, 2

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 33
	fmt.Println(j)
}