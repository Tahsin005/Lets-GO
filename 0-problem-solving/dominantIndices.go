func dominantIndices(nums []int) int {
    n, totalSum := len(nums), 0

    for _, v := range nums {
        totalSum += v
    }

    rightSum , count := totalSum, 0

    for i := 0; i < n - 1; i++ {
        rightSum -= nums[i]
        rightCount := n - i - 1
        if nums[i] * rightCount > rightSum {
            count++
        }
    }

    return count
}
