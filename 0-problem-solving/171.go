package main

import "fmt"

// https://leetcode.com/problems/find-if-digit-game-can-be-won/description/?envType=problem-list-v2&envId=array

func canAliceWin(nums []int) bool {
    singleDigit, doubleDigit := 0, 0
    for _, j := range (nums) {
        if j < 10 {
            singleDigit += j
        } else {
            doubleDigit += j
        }
    }
    return singleDigit != doubleDigit
}

func main () {

}