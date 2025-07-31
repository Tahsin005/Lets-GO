package main

import "fmt"

// https://leetcode.com/problems/sum-of-unique-elements/description/?envType=problem-list-v2&envId=hash-table

func sumOfUnique(nums []int) int {
	mapper := make(map[int]int)
	var sum int

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if v, ok := mapper[num]; !ok {
			sum += num
		} else if v == 1 {
			sum -= num

		}

		mapper[num]++
	}

	return sum
}

func main () {

}