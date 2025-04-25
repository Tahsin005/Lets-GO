package main

import "fmt"

// https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/?envType=problem-list-v2&envId=array

func nextGreatestLetter(letters []byte, target byte) byte {
    for _, char := range letters {
        if char > target {
            return char
        }
    }

    return letters[0]
}

func main () {

}