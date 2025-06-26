package main

import "fmt"

// https://leetcode.com/problems/find-the-number-of-good-pairs-i/description/?envType=problem-list-v2&envId=array

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
    count := 0
    for i := 0; i < len(nums1); i++ {
        for j := 0; j < len(nums2); j++ {
            if k * nums2[j] != 0 && nums1[i] % (k * nums2[j]) == 0 {
                count++
            }
        }
    }
    return count
}

func main () {

}