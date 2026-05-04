func countOppositeParity(nums []int) (r []int) {
	for i := 0; i < len(nums); i++ {
		c, searchOdd := 0, false
		if nums[i] % 2 == 0 {
			searchOdd = true
		}
		for j := i + 1; j < len(nums); j++ {
			if searchOdd {
				if nums[j] % 2 != 0 {
					c++
				}
			} else {
				if nums[j] % 2 == 0 {
					c++
				}
			}
		}
		r = append(r, c)
	}
	return
}
