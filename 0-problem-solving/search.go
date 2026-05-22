func search(nums []int, target int) int {
	bit0 := 0 
	if target >= nums[0] {
		bit0 = 1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}

		bit1 := 0
		if nums[mid] >= nums[0] {
			bit1 = 1
		}
		bit2 := 0
		if nums[mid] > target {
			bit2 = 1
		}

		mask := bit0 ^ bit1 ^ bit2
		if mask == 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return -1
}
