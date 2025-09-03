package main

import "fmt"

// https://leetcode.com/problems/reverse-degree-of-a-string/description/?envType=problem-list-v2&envId=n6ww7n9h

func abs(x int) int{
    if x < 0 {
        return -x
    }
    return x
}

func reverseDegree(s string) int {
    sum := 0
    j := 0

    for i := 0; i < len(s); i++ {
        char := rune(s[i])
		position := int(unicode.ToLower(char)) - int('z') - 1
        j = j + 1
        sum = sum + abs(position) * j
    }

    return sum
}

func main () {

}