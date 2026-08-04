func maxPairStrength(nums []int) (r int64) {
	for i := 0; i < len(nums) - 1; i++ {
		x := nums[i]
		for j := i + 1; j < len(nums); j++ {
			y := nums[j]
			z := (x * y) / (GCD(x, y) * GCD(x, y))
			if int64(z) > r {
				r = int64(z)
			}

		}
	}
	return
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}
