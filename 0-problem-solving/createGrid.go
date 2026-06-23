func createGrid(m int, n int) (r []string) {
	for i := 0; i < m; i++ {
		s := ""
		for j := 0; j < n; j++ {
			if i == 0 {
				s += "."
			} else {
				if j == n - 1 {
					s += "."
				} else {
					s += "#"
				}
			}
		}
		r = append(r, s)
	}
	return
}
