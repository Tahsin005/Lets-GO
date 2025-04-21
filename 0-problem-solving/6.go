package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/check-if-numbers-are-ascending-in-a-sentence/description/?envType=problem-list-v2&envId=string

func areNumbersAscending(s string) bool {
    words := strings.Split(s, " ")
    max := -1

    for _, v := range words {
        n, err := strconv.Atoi(v)

        if err != nil {
            continue
        }

        if n >= 1 && n < 100 {
            if n > max {
                max = n
            } else {
                return false
            }
        }
    }

    return true
}
func main () {
	fmt.Println(areNumbersAscending("1 box has 3 blue 4 red 6 green and 12 yellow marbles"))
	fmt.Println(areNumbersAscending("hello world 5 x 5"))
}