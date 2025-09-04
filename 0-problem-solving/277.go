package main

import "fmt"

// https://leetcode.com/problems/find-closest-person/description/?envType=daily-question&envId=2025-09-04

func findClosest(x int, y int, z int) int {
    person1, person2 := abs(x - z), abs(y - z)

    if person1 < person2 {
        return 1
    } else if person1 > person2 {
        return 2
    }
    return 0
}

func abs (x int) int {
    if x < 0 {
        return x * -1
    }
    return x
}

func main () {

}