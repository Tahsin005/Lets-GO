/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}

	prev := dummy
	curr := head

	for curr != nil {
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
		}

		if prev.Next == curr {
			prev = prev.Next
		} else {
			prev.Next = curr.Next
		}

		curr = curr.Next
	}

	return dummy.Next
}