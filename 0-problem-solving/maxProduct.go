func maxProduct(n int) int {
	numsArr := []int{}
	for n > 0 {
		numsArr = append(numsArr, n%10)
		n = n / 10
	}
	sort.Ints(numsArr)
	return numsArr[len(numsArr) - 2] * numsArr[len(numsArr) - 1]
}
