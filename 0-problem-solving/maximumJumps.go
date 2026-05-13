func maximumJumps(nums []int, target int) int {
    n := len(nums)
    dp := make([]int, n)
    for i := 1; i < n; i++ {
        dp[i] = -1
        for j := 0; j < i; j++ {
            if abs(nums[i] - nums[j]) <= target && dp[j] >= 0 {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
    }
    return dp[n - 1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
