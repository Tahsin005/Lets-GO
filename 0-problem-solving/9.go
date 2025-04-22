package main

import "fmt"

// https://leetcode.com/problems/add-to-array-form-of-integer/?envType=problem-list-v2&envId=array

func addToArrayForm(num []int, k int) []int {
    result := []int{}
    i := len(num) - 1

    for k > 0 || i >= 0 {
        if i >= 0 {
            k += num[i]
        }
        result = append(result, k % 10)
        k /= 10
        i--
    }

    for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
		result[i], result[j] = result[j], result[i]
	}

    return result
}


func main () {
	
}