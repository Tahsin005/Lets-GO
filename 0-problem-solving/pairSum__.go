/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    s := make([]int, 0)
    
    for head != nil {
        s = append(s, head.Val)
        head = head.Next
    }

    left := 0
    right := len(s) - 1
    max := 0

    for left < right {
        if s[left] + s[right] > max {
            max = s[left] + s[right]
        }
        left++
        right--
    }
    return max
}
