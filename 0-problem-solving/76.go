package main

import "fmt"

// https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/description/?envType=daily-question&envId=2025-05-27

func differenceOfSums(n, m int) (result int) {
    if m == 0 {
        return 0
    }

    return n * (n + 1) >> 1 - m * (n / m) * (n / m + 1)
}

func main () {

}