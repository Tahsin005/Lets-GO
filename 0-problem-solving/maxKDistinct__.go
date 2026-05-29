func maxKDistinct(nums []int, k int) []int {
    sort.Ints(nums)
    res := make([]int, 1, k)
    res[0] = nums[len(nums) - 1]
    for i := len(nums) - 2; i >= 0; i-- {
        if k > 1 && nums[i] != nums[i + 1] {
            res = append(res, nums[i])
            k--
        }
    }
    return res
}
