package main

import "fmt"

// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-i/description/?envType=daily-question&envId=2025-07-20

func maximumLength(nums []int) int {
    allEven, allOdd := 0, 0
    altEvenOdd, altOddEven := 0, 0

    for _, num := range nums {
        if num % 2 == 0 {
            allEven++
        } else {
            allOdd++
        }
    }

    expect := 0
    for _, num := range nums {
        if num % 2 == expect {
            altEvenOdd++
            expect ^= 1
        }
    }

    expect = 1
    for _, num := range nums {
        if num % 2 == expect {
            altOddEven++
            expect ^= 1
        }
    }

    max := allEven
    if allOdd > max {
        max = allOdd
    }
    if altEvenOdd > max {
        max = altEvenOdd
    }
    if altOddEven > max {
        max = altOddEven
    }

    return max
}

func main () {

}