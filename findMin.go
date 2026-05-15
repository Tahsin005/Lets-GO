func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for nums[right] < nums[left] {
		mid := left + (right - left) / 2

		if nums[left] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
