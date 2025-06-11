package main

import "fmt"

// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/description/?envType=daily-question&envId=2025-06-10

import (
    "math"
)

func maxDifference(s string) int {
    freq := make(map[rune]int)

    for _, ch := range s {
        freq[ch]++
    }

    oddMax := math.MinInt32
    evenMin := math.MaxInt32

    for _, count := range freq {
        if count % 2 == 1 {
            if count > oddMax {
                oddMax = count
            }
        } else {
            if count < evenMin {
                evenMin = count
            }
        }
    }

    return oddMax - evenMin
}

func main () {

}