package main

import "fmt"

// https://leetcode.com/problems/degree-of-an-array/description/?envType=problem-list-v2&envId=n6ww7n9h

func findShortestSubArray(nums []int) int {
	res := 1
	starts := make(map[int]int)
	counters, maxCounter := make(map[int]int), 1

	for i, num := range nums {
		if _, ok := starts[num]; ok {
			counters[num]++
			if counters[num] > maxCounter {
				maxCounter = counters[num]
				res = i - starts[num] + 1
			} else if counters[num] == maxCounter && i - starts[num] < res {
				res = i - starts[num] + 1
			}
		} else {
			starts[num], counters[num] = i, 1
		}
	}
	return res
}

func main () {

}