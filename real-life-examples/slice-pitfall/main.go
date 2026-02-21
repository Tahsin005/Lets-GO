package main

import "fmt"

func main() {
	// repeated append without pre-alloc (O(N) copies)

	var a []int // len=0, cap=0

	for i := 0; i < 8; i++ {
		a = append(a, i)
		// when cap is exceeded:
		//  - new array allocated
		//  - all existing elements copied (O(n))
		fmt.Printf("append %d -> len=%d cap=%d\n", i, len(a), cap(a))
	}

	fmt.Println()

	// pre-allocating capacity avoids reallocations

	b := make([]int, 0, 8) // capacity known in advance

	for i := 0; i < 8; i++ {
		b = append(b, i)
		// no reallocation, no copying
		fmt.Printf("append %d -> len=%d cap=%d\n", i, len(b), cap(b))
	}

	fmt.Println()

	// slice aliasing (shared backing array)

	x := make([]int, 0, 4)
	x = append(x, 1, 2)

	y := x               // y shares the same backing array
	y = append(y, 3)     // capacity not exceeded → mutates x too

	fmt.Println("x:", x) // [1 2 3]  ← unexpected
	fmt.Println("y:", y)

	fmt.Println()

	// force a copy to avoid shared backing array

	p := []int{1, 2}
	q := append([]int(nil), p...) // deep copy

	q = append(q, 3)

	fmt.Println("p:", p) // [1 2]
	fmt.Println("q:", q) // [1 2 3]

	fmt.Println()

	// capacity cliff (latency spike)

	c := make([]int, 0, 1)

	for i := 0; i < 10; i++ {
		oldCap := cap(c)
		c = append(c, i)
		if cap(c) != oldCap {
			// this append caused:
			//  - allocation
			//  - full copy of old elements (O(n))
			fmt.Printf("CAPACITY GROWTH at i=%d: %d -> %d\n",
				i, oldCap, cap(c))
		}
	}

	fmt.Println()

	// assuming append is always O(1)

	// amortized O(1), but individual append can be O(n)
	// this matters in hot loops, queues, buffers, etc.

	fmt.Println("Slices are cheap until capacity grows.")
}