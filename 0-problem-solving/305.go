package main

import "fmt"

// https://leetcode.com/problems/count-asterisks/description/

func countAsterisks(s string) int {
    count := 0
    withinPair := false

    for _, char := range s {
        if char == '|' {
            withinPair = !withinPair
        } else if char == '*' && !withinPair {
            count++
        }
    }

    return count
}

func main () {

}