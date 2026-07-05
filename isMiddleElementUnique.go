func isMiddleElementUnique(nums []int) bool {
	n := len(nums)
	mid := n / 2
	midVal := nums[mid]
	cnt := 0

	for _, num := range nums {
		if num == midVal {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}

	return true
}
