package main

import "fmt"

// https://leetcode.com/problems/sum-multiples/description/

func sumOfMultiples(n int) (sum int) {
    for i := 1; i <= n; i++ {
        if i % 5 == 0 || i % 3 == 0 || i % 7 == 0 {
            sum += i
        }
    }

    return
}

func main () {

}