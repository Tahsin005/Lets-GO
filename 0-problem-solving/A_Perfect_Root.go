package main

import "fmt"

func main() {
	var tt int 
	fmt.Scan(&tt)

	for tt > 0 {
		var n int 
		fmt.Scan(&n)

		for i := 0; i < n; i++ {
			fmt.Printf("%v ", i + 1)
		}
		fmt.Println()
		tt--
	}
}