func smallestRangeI(nums []int, k int) int {
	min := nums[0]
	max := nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	result := (max - k) - (min + k)
	if result < 0 {
		return 0
	}
	return result

}
