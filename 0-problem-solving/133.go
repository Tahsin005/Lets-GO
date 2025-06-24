package main

import "fmt"

// https://leetcode.com/problems/minimum-average-of-smallest-and-largest-elements/description/?envType=problem-list-v2&envId=array

func minimumAverage(nums []int) float64 {
    sort.Ints(nums)
    ans := 50.0
    i := 0
    for {
        minVal := nums[0]
        maxVal := nums[len(nums) - 1]

        nums = nums[1:]
        nums = nums[:len(nums) - 1]

        avg := float64((float64(minVal) + float64(maxVal)) / 2)

        if avg < ans {
            ans = avg
        }

        i++

        if len(nums) == 0 {
            break
        }
    }

    return ans
}

func main() {
	
}