package main

import "fmt"

// https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/description/?envType=problem-list-v2&envId=array

func mostWordsFound(sentences []string) int {
    var ans int
    for _, sentence := range sentences {
        count := 0
        for _, letter := range sentence {
            if letter == ' ' {
                count++
            }
        }
        if count > ans {
            ans = count
        }
    }

    return ans + 1
}

func main () {

}