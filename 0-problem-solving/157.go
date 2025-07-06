package main

import "fmt"

// https://leetcode.com/problems/find-the-original-typed-string-i/description/?envType=daily-question&envId=2025-07-06

func possibleStringCount(word string) int {
    n := len(word)
    count := 1
    i := 0
    for i < n {
        j := i
        for j < n && word[j] == word[i] {
            j++
        }
        length := j - i
        if length > 1 {
            count += (length - 1)
        }
        i = j
    }
    return count
}

func main () {

}