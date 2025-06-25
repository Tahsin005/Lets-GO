package main

import "fmt"

// https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/description/?envType=problem-list-v2&envId=array

func minimumOperations(nums []int) int {
    count := 0
    for i := 0; i < len(nums); i++ {
        if (nums[i] - 1) % 3 == 0 || (nums[i] + 1) % 3 == 0 {
            count++
        }
    }
    return count
}
func main () {

}