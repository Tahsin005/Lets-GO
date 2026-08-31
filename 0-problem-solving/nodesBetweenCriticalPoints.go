/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isCritical(a, x, b int) bool {
	return (a > x && x < b) || (a < x && x > b)
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	impossible := []int{-1, -1}

	var points []int

	for head != nil {
		points = append(points, head.Val)
		head = head.Next
	}

	if len(points) <= 3 {
		return impossible
	}

	var idxs []int

	for i := 1; i < len(points) - 1; i++ {
		if isCritical(points[i - 1], points[i], points[i + 1]) {
			idxs = append(idxs, i+1)
		}
	}

	if len(idxs) < 2 {
		return impossible
	}

	mn := int(^uint(0) >> 1)

	for i := 0; i < len(idxs) - 1; i++ {
		diff := idxs[i + 1] - idxs[i]
		if diff < mn {
			mn = diff
		}
	}

	mx := idxs[len(idxs) - 1] - idxs[0]

	return []int{mn, mx}
}
