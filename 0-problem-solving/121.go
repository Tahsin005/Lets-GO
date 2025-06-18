package main

import (
	"strconv"
)

// https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/description/?envType=daily-question&envId=2025-06-18

func minMaxDifference(num int) int {
    s := strconv.Itoa(num)
    var ch rune
    for _, c := range s {
        if c != '9' {
            ch = c
            break
        }
    }
    mx := ""
    if ch != 0 {
        for _, c := range s {
            if c == ch {
                mx += "9"
            } else {
                mx += string(c)
            }
        }
    } else {
        mx = s
    }
    ch0 := rune(s[0])
    mn := ""
    for _, c := range s {
        if c == ch0 {
            mn += "0"
        } else {
            mn += string(c)
        }
    }
    mxVal, _ := strconv.Atoi(mx)
    mnVal, _ := strconv.Atoi(mn)
    return mxVal - mnVal
}

func main() {

}