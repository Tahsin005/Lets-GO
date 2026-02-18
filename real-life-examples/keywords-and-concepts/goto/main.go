package main

import "fmt"

func main() {
	i := 0

Start:
	if i < 3 {
		fmt.Println(i)
		i++
		goto Start
	}
}
