func constructProductMatrix(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])
	modulus := 12345

	// suffixProduct[r][c]:= the product % modulus of numbers after grid[r][c]
	suffixProduct := make([][]int, m)
	accProduct := 1
	for r := m - 1; r >= 0; r-- {
		suffixProduct[r] = make([]int, n)
		for c := n - 1; c >= 0; c-- {
			suffixProduct[r][c] = accProduct
			accProduct = accProduct * grid[r][c] % modulus
		}
	}

	productMatrix := make([][]int, m)
	prefixProduct := 1
	for r := range m {
		productMatrix[r] = make([]int, n)
		for c := range n {
			productMatrix[r][c] = prefixProduct * suffixProduct[r][c] % modulus
			prefixProduct = prefixProduct * grid[r][c] % modulus
		}
	}

	return productMatrix
}
