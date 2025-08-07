package main

import "fmt"

// https://leetcode.com/problems/check-if-any-element-has-prime-frequency/description/

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= n / 2; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func checkPrimeFrequency(nums []int) bool {
	var m map[int]int = make(map[int]int)
	for _, a := range nums {
		m[a]++
	}
	for i := 0; i <= 100; i++ {
		v, exists := m[i]
		if exists {
			if isPrime(v) {
				return true
			}
		}
	}
	return false
}

func main () {

}