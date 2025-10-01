package main

import "fmt"

// https://leetcode.com/problems/xor-operation-in-an-array/

func xorOperation(n int, start int) int {
	res := 0

	for i := 0; i < n; i++ {
		res ^= start + 2 * i
	}

	return res
}

func main () {

}