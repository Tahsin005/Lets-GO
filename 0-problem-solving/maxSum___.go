func maxSum(nums []int, k int, mul int) (r int64) {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if k > 0 {
			if mul > 0 {
				r += int64(nums[i] * mul)
				mul--
			} else {
				r += int64(nums[i])
			}
			k--
		} else {
			break
		}
	}
	return
}
