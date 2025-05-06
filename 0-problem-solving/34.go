package main

import "fmt"

// https://leetcode.com/problems/relative-sort-array/description/?envType=problem-list-v2&envId=array

func relativeSortArray(arr1 []int, arr2 []int) []int {
    pos := make(map[int]int)
    for idx, val := range arr2 {
        pos[val] = idx
    }

    sort.Slice(arr1, func(i, j int) bool {
        idx1, ok1 := pos[arr1[i]]
        idx2, ok2 := pos[arr1[j]]

        if ok1 && ok2 {
            return idx1 < idx2
        } else if ok1 {
			return true
		} else if ok2 {
			return false
		} else {
			return arr1[i] < arr1[j]
		}
    })

    return arr1
}

func main () {

}