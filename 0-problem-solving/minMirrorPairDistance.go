func minMirrorPairDistance(nums []int) int {
	n := len(nums)
	
	mp := make(map[int]int)
	ans := len(nums) + 1
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		rv := reverse(num)
		if val, ok := mp[rv]; ok {
			ans = min(ans, val - i)
		}

		mp[num] = i
	}

	if ans == len(nums) + 1 {
		return -1
	}
	return ans
}

func reverse(num int) int {
	x := 0
	for num > 0 {
		x = x * 10 + num % 10
		num /= 10
	}
	return x
}
