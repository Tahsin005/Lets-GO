package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/hexadecimal-and-hexatrigesimal-conversion/description/

func concatHex36(n int) string {
    return base16(n * n) + base36(n * n * n)
}

func base16(n int) string {
    out := ""
    for n > 0 {
        if n % 16 > 9 {
            out = string('A' + (n % 16 - 10)) + out
        } else {
            out = strconv.Itoa(n % 16) + out
        }
        n /= 16
    }
    return out
}

func base36(n int) string {
    out := ""
    for n > 0 {
        if n % 36 > 9 {
            out = string('A' + (n % 36 - 10)) + out
        } else {
            out = strconv.Itoa(n % 36) + out
        }
        n /= 36
    }
    return out
}

func main () {

}