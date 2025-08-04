package main

import "fmt"

// https://leetcode.com/problems/split-array-by-prime-indices/description/

func splitArray(nums []int) int64 {
    n := len(nums)
    if n == 1 {
        return int64(nums[0])
    }
    prime := make([]bool, n)
    prime[0] = true
    prime[1] = true
    for i := 2; i <= int(math.Sqrt(float64(n + 1))); i++ {
        for j := i + i; j < n; j = j + i {
            prime[j] = true
        }
    }

    sum1, sum2 := int64(0), int64(0)
    for i := 0; i < n; i++ {
        if prime[i] {
            sum1 += int64(nums[i])
        } else {
            sum2 += int64(nums[i])
        }
    }
    return Abs(sum2 - sum1)
}
func Abs(a int64)int64 {
    if a < 0 {
        return -a
    }
    return a
}

func main () {

}