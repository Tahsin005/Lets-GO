func solve(num int, dig int) int {
	ans := 0
	for num > 0 {
		d := num % 10
		num /= 10
		if d == dig {
			ans++
		}
	}
	return ans
}

func countDigitOccurrences(nums []int, digit int) int {
	ans := 0

	for _, x := range nums {
		ans += solve(x, digit)
	}

	return ans
}
