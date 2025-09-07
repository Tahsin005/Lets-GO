package main

import "fmt"

// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/description/?envType=daily-question&envId=2025-09-07

func sumZero(n int) []int {
    result := []int{}

    for i := 1; i <= n / 2; i++ {
        result = append(result, -i)
        result = append(result, i)
    }

    if n % 2 != 0 {
        result = append(result, 0)
    }

    return result
}

func main () {

}