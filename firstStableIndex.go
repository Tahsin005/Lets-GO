func firstStableIndex(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		left, right := nums[: i + 1], nums[i:]
		h, l := 0, 0
		for j := 0; j < len(left); j++ {
			if j == 0 {
				h = left[j]
			}
			if left[j] > h {
				h = left[j]
			}
		}
		for j := 0; j < len(right); j++ {
			if j == 0 {
				l = right[j]
			}
			if right[j] < l {
				l = right[j]
			}
		}
		if h - l <= k {
			return i
		}
	}
	return -1
}
