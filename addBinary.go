func addBinary(a string, b string) string {
	x, y := new(big.Int), new(big.Int)
	x, _ = x.SetString(a, 2)
	y, _ = y.SetString(b, 2)

	sum := new(big.Int)
	sum.Add(x, y)

	return sum.Text(2)
}
