package main

import "fmt"

// https://leetcode.com/problems/check-if-the-sentence-is-pangram/description/

func checkIfPangram(sentence string) bool {
    if len(sentence) < 26 {
        return false
    }

    hash := make([]int, 26)
    for i := 0; i < len(sentence); i++ {
        ind := sentence[i] - 'a'
        hash[ind]++
    }

    for i := 0; i < 26; i++ {
        if hash[i] == 0 {
            return false
        }
    }

    return true
}

func main () {

}