func longestBalanced(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	mx := nums[0]
	for _, v := range nums {
		if v > mx {
			mx = v
		}
	}

	ans := 0
	count := make([]int, mx+1)
	last := make([]int, mx+1)
	next := make([]int, n)
	uniq := [2]int{}

	for i := range last {
		last[i] = n
	}
	for i := range next {
		next[i] = n
	}

	for i, x := range nums {
		update(x, 1, count, &uniq)
		if last[x] < n {
			next[last[x]] = i
		}
		last[x] = i
	}

	j := n
	for i := 0; i < n-1; i++ {
		if n-i <= ans {
			break
		}

		for j-i > ans {
			if uniq[0] == uniq[1] {
				ans = j - i
				break
			}
			j--
			update(nums[j], -1, count, &uniq)
		}

		for j < next[i] {
			update(nums[j], 1, count, &uniq)
			j++
		}

		update(nums[i], -1, count, &uniq)
	}

	return ans
}

func update(x int, val int, count []int, uniq *[2]int) {
	count[x] += val

	if count[x] == 0 && val < 0 {
		uniq[x&1]--
	} else if count[x] == val && val > 0 {
		uniq[x&1]++
	}
}
