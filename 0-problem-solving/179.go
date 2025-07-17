package main

import "fmt"

// https://leetcode.com/problems/remove-outermost-parentheses/description/

func removeOuterParentheses(s string) string {
    var result string
    var index int

    for _, x := range s {
        switch x {
        case '(':
            if index > 0 {
                result += string(x)
            }
            index++
        case ')':
            index--
            if index > 0 {
                result += string(x)
            }
        }
    }

    return result
}

func main () {

}