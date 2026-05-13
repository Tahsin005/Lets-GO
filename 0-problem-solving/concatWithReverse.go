func concatWithReverse(nums []int) []int {
    l := len(nums)
    res := make([]int, l * 2)
    for i, j := 0, l - 1; i < l; i, j = i + 1, j - 1 {
        res[i], res[l + i] = nums[i], nums[j]
    } 
    return res
}
