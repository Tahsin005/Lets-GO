func countValidSubarrays(nums []int, x int) (r int) {
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			last := sum % 10
			first := sum
			for first >= 10 {
				first /= 10
			}
			if first == x && last == x {
				r++
			}
		}
	}
	return
}
