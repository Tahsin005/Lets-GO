package main

import "fmt"

// https://leetcode.com/problems/minimum-element-after-replacement-with-digit-sum/description/?envType=problem-list-v2&envId=array

func minElement(nums []int) int {
    res := 1 << 31 - 1

    for _,v := range(nums) {
        res = min(digitSum(v), res)
    }

    return res
}

func digitSum(n int) int {
    res := 0

    for n > 0 {
        res += n % 10
        n /= 10
    }

    return res
}
func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
	fmt.Println(123)
}