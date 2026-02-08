package main

import (
	"fmt"
	// "time"
)

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func sum(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	c <- total
}

func main() {
	// go say("world")
	// say("hello")

	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x := <-c // receiving from channel is a blocking execution
	y := <-c

	fmt.Println(x, y, x+y)
}