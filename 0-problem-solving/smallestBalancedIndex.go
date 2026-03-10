func smallestBalancedIndex(nums []int) int {
	sum := uint64(0)
	for _, n := range nums {
		sum += uint64(n)
	}

	prod, result := uint64(1), -1
	for i, hi := len(nums) - 1, uint64(0); 0 <= i; i-- {
		sum -= uint64(nums[i])
		if sum == prod {
			result = i
		}
		hi, prod = bits.Mul64(prod, uint64(nums[i]))
		if 0 < hi {
			return result
		}
	}

	return result
}
