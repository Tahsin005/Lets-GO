package main

import "fmt"

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/description/?envType=daily-question&envId=2025-07-15

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    res := 0

    for head != nil {
        res <<= 1
        res |= head.Val
        head = head.Next
    }

    return res
}

func main () {

}