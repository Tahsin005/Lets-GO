package main

import "fmt"

// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/description/?envType=problem-list-v2&envId=array

func sumZero(n int) []int {
    answer := make([]int, n)

    x := n / 2
    if n % 2 == 0 {
        answer[n - 1] = 0
    }
    for i := 1; i <= x; i++ {
        answer[i - 1] = -i
        answer[i - 1 + x] = i
    }
    return answer
}

func main () {

}