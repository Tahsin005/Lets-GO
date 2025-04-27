package main

import (
	"fmt"
	"time"
)

func concurrrentFib(n int) {
	chInts := make(chan int)
	go fibonacci(n, chInts)

	for item := range chInts {
		fmt.Printf("Received %v\n", item)
	}
}

func test(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x + y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}
