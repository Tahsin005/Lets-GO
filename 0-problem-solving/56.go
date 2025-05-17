package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/description/?envType=problem-list-v2&envId=array

func abs(a int) int {
    if a < 0 {
        a = -a
    }
    return a
}

func minMovesToSeat(seats []int, students []int) int {
    ans := 0
    sort.Ints(seats)
    sort.Ints(students)
    for i := 0; i < len(seats); i++ {
        ans += abs(seats[i] - students[i])
    }
    return ans
}

func main () {

}