func missingInteger(nums []int) int {
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i - 1] + 1 {
			break
		}
		sum += nums[i]
	}

	sort.Ints(nums)
	for i := range nums {
		if nums[i] == sum {
			sum++
		}
	}

	return sum
}
