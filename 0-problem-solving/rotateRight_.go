/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {

	count := 1
	currentnode := head
	if head == nil || head.Next == nil {
		return head
	}
	for currentnode.Next != nil {
		currentnode = currentnode.Next
		count++
	}
	n := count
	k = k % n
	if k == 0 {
		return head
	}
	steps := n - k
	node := head
	for i := 1; i < steps; i++ {
		node = node.Next
	}
	newHead := node.Next
	currentnode.Next = head
	node.Next = nil

	return newHead

}
