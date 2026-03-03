func minDistinctFreqPair(nums []int) []int {
    f := make(map[int]int)
    for _, n := range nums {
        f[n]++
    }
    x, y := -1, -1
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if nums[i] == nums[j] {
                continue
            }
            if f[nums[i]] == f[nums[j]] {
                continue
            }
            if nums[i] >= nums[j] {
                continue
            }
            if x == -1 || x > nums[i] || (x == nums[i] && y > nums[j]) {
                x = nums[i]
                y = nums[j]
            }
        }
    }
    return []int{x, y}
}
