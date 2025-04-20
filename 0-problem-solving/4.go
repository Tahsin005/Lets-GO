package main

import "fmt"

// https://leetcode.com/problems/minimum-time-to-type-word-using-special-typewriter/description/?envType=problem-list-v2&envId=string

func min (a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs (a int) int {
    if a < 0 {
        return a * -1
    }
    return a
}
func minTimeToType(word string) int {
    curr := 0
    seconds := 0

    for _, char := range word {
        c := int(char - 97)
        seconds += min(abs(c - curr), abs(abs(c - curr) - 26)) + 1
        curr = c
    }

    return seconds
}


func main () {
	fmt.Println("tahsin.ferdous3546@gmail.com")
}