package main

import "fmt"

// https://leetcode.com/problems/find-common-elements-between-two-arrays/description/?envType=problem-list-v2&envId=array

func findIntersectionValues(nums1 []int, nums2 []int) []int {

    m1 := map[int]bool{}
    m2 := map[int]bool{}

    for _, v := range(nums1) {
        m1[v] = true
    }
    for _, v := range(nums2) {
        m2[v] = true
    }
    c1 := 0
    for _, v := range(nums1) {
        if m2[v] {
            c1++
        }
    }
    c2 := 0
    for _, v := range(nums2) {
        if m1[v]{
            c2++
        }
    }
    return []int{c1,c2}
}

func main () {

}