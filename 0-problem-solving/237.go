package main

import "fmt"

// https://leetcode.com/problems/power-of-four/?envType=daily-question&envId=2025-08-15

func isPowerOfFour(n int) bool {
    if n <= 0 {
		return false
	}
	for n % 4 == 0 {
		n /= 4
	}
	return n == 1
}

func main () {

}