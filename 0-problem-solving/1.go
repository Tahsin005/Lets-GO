package main

import "fmt"

// https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/description/?envType=problem-list-v2&envId=string

func checkString(s string) bool {
    seenB := true

    for _, char := range s {
        if char == 'b' {
            seenB = true
        }
        if char == 'a' && seenB {
            return false
        }
    }

    return true
}

func main () {
	fmt.Println("Given a string s consisting of only the characters 'a' and 'b', return true if every 'a' appears before every 'b' in the string. Otherwise, return false.")
    fmt.Println(checkString("aaabbb")) // true
    fmt.Println(checkString("abab"))   // false
    fmt.Println(checkString("bbbb")) // true
    fmt.Println(checkString("aaaa")) // true
}