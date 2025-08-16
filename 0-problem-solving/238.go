package main

import "fmt"

// https://leetcode.com/problems/maximum-69-number/?envType=daily-question&envId=2025-08-16

func maximum69Number(num int) int {
    s := []byte(strconv.Itoa(num))
    for i := 0; i < len(s); i++ {
        if s[i] == '6' {
            s[i] = '9'
            break
        }
    }
    res, _ := strconv.Atoi(string(s))
    return res
}

func main () {

}