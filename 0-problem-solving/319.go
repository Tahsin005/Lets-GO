package main

import "fmt"

// https://leetcode.com/problems/maximum-total-subarray-value-i/description/

func maxTotalValue(nums []int, k int) int64 {
    var result int64

    maxNum, minNum := getMaxAndMin(nums)
    result = (maxNum - minNum) * int64(k)

    return result
}

func getMaxAndMin(nums []int) (int64, int64) {
    var maxNum, minNum int64 = math.MinInt, math.MaxInt

    for _, num := range nums {
        maxNum = max(maxNum, int64(num))
        minNum = min(minNum, int64(num))
    }

    return maxNum, minNum
}

func main () {

}