package main

import "fmt"

// https://leetcode.com/problems/find-the-integer-added-to-array-i/description/?envType=problem-list-v2&envId=array

func addedInteger(nums1 []int, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    diff := nums2[0] - nums1[0]
    for i := 0; i < len(nums1); i++ {
        if nums2[i] - nums1[i] != diff {
            break
            return -1
        } else {
            continue
        }
    }
    return diff
}

func main() {
    
}