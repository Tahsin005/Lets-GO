func aggregateTimeSeries(series1 [][]int, series2 [][]int) [][]int {
	ans := [][]int{}
	n1, n2 := len(series1), len(series2)
	i, j := 0, 0

	for i < n1 || j < n2 {
		var t int

		if i == n1 {
			t = series2[j][0]
		} else if j == n2 {
			t = series1[i][0]
		} else {
			if series1[i][0] < series2[j][0] {
				t = series1[i][0]
			} else {
				t = series2[j][0]
			}
		}

		x1 := 0
		if i < n1 {
			x1 = series1[i][1]
		}

		x2 := 0
		if j < n2 {
			x2 = series2[j][1]
		}

		ans = append(ans, []int{t, x1 + x2})

		if i < n1 && series1[i][0] == t {
			i++
		}

		if j < n2 && series2[j][0] == t {
			j++
		}
	}

	return ans
}
