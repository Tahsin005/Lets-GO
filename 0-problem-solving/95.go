package main

import "fmt"

// https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing/description/?envType=problem-list-v2&envId=array

func minOperations(nums []int) (r int) {
    for i := 0; i < len(nums); i++{
        if i - 1 >= 0 {
            for nums[i] <= nums[i - 1] {
                nums[i]++
                r++
            }
        }
    }
	return
}

func main () {

}