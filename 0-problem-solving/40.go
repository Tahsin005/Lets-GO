package main

import "fmt"

// https://leetcode.com/problems/decompress-run-length-encoded-list/description/?envType=problem-list-v2&envId=array

func decompressRLElist(nums []int) []int {
    var res []int

    for i := 0; i < len(nums); i += 2 {
        for count := 0; count < nums[i]; count++ {
            res = append(res, nums[i + 1])
        }
    }

    return res
}

func main () {

}