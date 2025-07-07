package main

import "fmt"

// https://leetcode.com/problems/find-the-k-th-character-in-string-game-i/description/?envType=daily-question&envId=2025-07-06

func kthCharacter(k int) byte {
    shift := 0
    k--
    for ; k > 0; k /= 2 {
        if k % 2 == 1 {
            shift++
        }
    }
    shift = shift % 26
    return byte('a') + byte(shift)
}

func main () {

}