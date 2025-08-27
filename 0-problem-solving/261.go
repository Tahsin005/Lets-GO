package main

import "fmt"

// https://leetcode.com/problems/gcd-of-odd-and-even-sums/description/

func gcdOfOddEvenSums(n int) int {
	var sumOdd int
	var sumEven int
	for i := 1; i <= n * 2; i++ {
		switch {
		case i % 2 == 0:
			sumEven += i
		default:
			sumOdd += i
		}
	}

    a := sumEven
    b := sumOdd

    for b != 0 {
        a, b = b, a % b
    }

	return a
}

func main () {

}