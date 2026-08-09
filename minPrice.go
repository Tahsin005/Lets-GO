func minPrice(prices []int, discounts []int) float64 {
	sort.Slice(prices, func(i, j int) bool {
		return prices[i] > prices[j]
	})

	sort.Slice(discounts, func(i, j int) bool {
		return discounts[i] > discounts[j]
	})

	var ans float64

	for i := 0; i < len(prices); i++ {
		if i < len(discounts) {
			ans += float64(prices[i]) * float64(100 - discounts[i]) / 100.0
		} else {
			ans += float64(prices[i])
		}
	}

	return ans
}
