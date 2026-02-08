package main

import (
	"fmt"
	"time"
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

func fibonacci(n int, c chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}
	close(c)
}

// example of select
func fibo(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
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

	fmt.Println("----------------")

	cc := make(chan int, 10)
	go fibonacci(cap(cc), cc)
	for i := range cc {
		fmt.Println(i)
	}

	fmt.Println("----------------")

	ccc := make(chan int)
	quit := make(chan int)
	go func ()  {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ccc)
		}
		quit <- 0
	}()
	fibo(ccc, quit)

	fmt.Println("----------------")

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}