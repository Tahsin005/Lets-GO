package main

import "fmt"

// https://leetcode.com/problems/fair-candy-swap/description/?envType=problem-list-v2&envId=array

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumA := 0
	sumB := 0
	for _, a := range aliceSizes {
		sumA += a
	}
	for _, b := range bobSizes {
		sumB += b
	}

	diff := (sumB - sumA) / 2

	setB := make(map[int]bool)
	for _, b := range bobSizes {
		setB[b] = true
	}

	for _, a := range aliceSizes {
		if _, exists := setB[a + diff]; exists {
			return []int{a, a + diff}
		}
	}

	return []int{}
}

func main () {

}