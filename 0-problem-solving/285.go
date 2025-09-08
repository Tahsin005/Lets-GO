package main

import "fmt"

// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/submissions/1763218854/?envType=daily-question&envId=2025-09-08

func getNoZeroIntegers(n int) []int {
    for i := 1; i < n; i++ {
        if zeros(i) == true && zeros(n - i) == true {
            return []int{i, n - i}
        }
    }
    return []int{}
}

func zeros(a int)bool {
    for a > 0 {
        if a % 10 == 0 {
            return false
        }
        a /= 10
    }
    return true
}

func main () {

}