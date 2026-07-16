func minSubsequence(nums []int) []int {
    sort.Ints(nums)
    idx := -1
    for i := len(nums)-1; i >= 0; i-- {
        if sum(nums[:i]) < sum(nums[i:]) {
            idx = i
            break
        }
    }
    res := slices.Clone(nums[idx:])
    slices.Reverse(res)
    return res
}

func sum(nums []int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}
