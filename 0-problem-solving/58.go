package main

import "fmt"

// https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/description/?envType=problem-list-v2&envId=array

func countKDifference(nums []int, k int) int {
	hash := make(map[int]int)

	for _, v := range nums {
		hash[v]++
	}

	counter := 0
	for _, v := range nums {
		counter += hash[v - k]
	}

	return counter
}

func main () {

}