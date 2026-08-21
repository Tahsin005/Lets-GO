func maximizeSum(nums []int, k int) int {
    sort.Ints(nums)
    sum := nums[len(nums) - 1]
    start := sum + 1
    last := sum + k
    for start < last {
        sum += start
        start++
    }
    return sum
}
