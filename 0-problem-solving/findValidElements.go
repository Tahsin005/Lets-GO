func findValidElements(nums []int) (r []int) {
	for i := 0; i < len(nums); i++ {
		if i == 0 || i == len(nums) - 1 {
			r = append(r, nums[i])
			continue
		}
		maxL, maxR := 0, 0
		for j := 0; j < i; j++ {
			if nums[j] > maxL {
				maxL = nums[j]
			}
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > maxR {
				maxR = nums[j]
			}
		}
		if nums[i] > maxL || nums[i] > maxR {
			r = append(r, nums[i])
		}
	}
	return
}
