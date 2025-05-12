package main

import (
	"math/bits"
	"sort"
)

// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/description/?envType=problem-list-v2&envId=array

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		countI := bits.OnesCount(uint(arr[i]))
		countJ := bits.OnesCount(uint(arr[j]))
		return countI < countJ || (countI == countJ && arr[i] < arr[j])
	})
	return arr
}

func main () {

}