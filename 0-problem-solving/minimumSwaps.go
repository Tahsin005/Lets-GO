func minimumSwaps(nums []int) int {
	zeros := 0
	swaps := 0
	n := len(nums)
	for _, num := range nums {
		if num == 0 {
			zeros++
			if nums[n-zeros] != 0 {
				swaps++
			}
		}
	}
	return swaps
}
