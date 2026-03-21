func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func gcdSum(nums []int) int64 {
	n := len(nums)

	maxi := nums[0]
	pref := make([]int, n)

	for i := 0; i < n; i++ {
		if nums[i] > maxi {
			maxi = nums[i]
		}
		pref[i] = gcd(maxi, nums[i])
	}

	sort.Ints(pref)

	l, r := 0, n - 1
	var ans int64 = 0

	for l < r {
		ans += int64(gcd(pref[l], pref[r]))
		l++
		r--
	}

	return ans
}
