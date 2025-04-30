package main

import "fmt"

// https://leetcode.com/problems/count-symmetric-integers/description/?envType=daily-question&envId=2025-04-25

func isSymmetrical(x int) bool {
    s := strconv.Itoa(x)
    if len(s) & 1 == 1 {
        return false
    }

    sumL, sumR := 0, 0
	mid := len(s) / 2

	for i := 0; i < mid; i++ {
		sumL += int(s[i] - '0')
		sumR += int(s[i + mid] - '0')
	}

	return sumL == sumR
}

func countSymmetricIntegers(low int, high int) int {
    cnt := 0

    for i := low; i <= high; i++ {
        if isSymmetrical(i) {
            cnt++
        }
    }

    return cnt
}

func main () {

}