import (
	"sync"
)

func longestBalanced(s string) int {
	var wg sync.WaitGroup
	results := make(chan int, 5)

	wg.Add(5)

	go func() {
		defer wg.Done()
		results <- calc1(s)
	}()

	go func() {
		defer wg.Done()
		results <- calc2(s, 'a', 'b')
	}()

	go func() {
		defer wg.Done()
		results <- calc2(s, 'b', 'c')
	}()

	go func() {
		defer wg.Done()
		results <- calc2(s, 'a', 'c')
	}()

	go func() {
		defer wg.Done()
		results <- calc3(s)
	}()

	wg.Wait()
	close(results)

	maxVal := 0
	for v := range results {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func calc1(s string) int {
	res := 0
	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		if j - i > res {
			res = j - i
		}
		i = j
	}
	return res
}

func calc2(s string, a, b byte) int {
	res := 0
	n := len(s)
	i := 0

	for i < n {
		for i < n && s[i] != a && s[i] != b {
			i++
		}

		mp := map[int]int{}
		mp[0] = i - 1
		diff := 0

		for i < n && (s[i] == a || s[i] == b) {
			if s[i] == a {
				diff++
			} else {
				diff--
			}

			if idx, ok := mp[diff]; ok {
				if i - idx > res {
					res = i - idx
				}
			} else {
				mp[diff] = i
			}
			i++
		}
	}
	return res
}

func key(x, y int) int64 {
	return (int64(x+100000) << 20) | int64(y+100000)
}

func calc3(s string) int {
	mp := map[int64]int{}
	mp[key(0, 0)] = -1

	cnt := [3]int{}
	res := 0

	for i := 0; i < len(s); i++ {
		cnt[s[i] - 'a']++

		x := cnt[0] - cnt[1]
		y := cnt[1] - cnt[2]
		k := key(x, y)

		if idx, ok := mp[k]; ok {
			if i - idx > res {
				res = i - idx
			}
		} else {
			mp[k] = i
		}
	}
	return res
}
