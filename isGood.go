func isGood(nums []int) bool {
    n := slices.Max(nums)
    if len(nums) != n + 1 {
        return false
    }
    arr := make([]int, n + 1)
    for _, d := range nums {
        arr[d]++
        if d != n && arr[d] > 1 {
            return false
        }
    }
    return arr[n] == 2
}
