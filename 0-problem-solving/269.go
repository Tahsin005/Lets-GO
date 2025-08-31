package main

import "fmt"

// https://leetcode.com/problems/find-the-least-frequent-digit/description/

func getLeastFrequentDigit(n int) int {
	f := make([]int, 10)

	for n > 0 {
		d := n % 10
		f[d]++
		n /= 10
	}

	ans := -1
	minu := math.MaxInt32
	for d := 0; d <= 9; d++ {
		if f[d] > 0 {
			if f[d] < minu || (f[d] == minu && (ans == -1 || d < ans)) {
				minu = f[d]
				ans = d
			}
		}
	}

	return ans
}

func main () {

}