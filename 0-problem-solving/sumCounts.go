func sumCounts(nums []int) (c int) {
    for i := range nums {
        set := [101]bool{}
        k := 0
        for j := i; j < len(nums); j++ {
            if !set[nums[j]] {
                set[nums[j]] = true
                k++
            }
            c += k * k
        }
    }
    return c
}
