func minimumDeletions(nums []int) int {
    n := len(nums)
    if n <= 2 {
        return n
    }

    minIdx, maxIdx := 0, 0
    for i := 1; i < n; i++ {
        if nums[i] < nums[minIdx] {
            minIdx = i
        }
        if nums[i] > nums[maxIdx] {
            maxIdx = i
        }
    }

    mini, maxi := minIdx, maxIdx
    if mini > maxi {
        mini, maxi = maxi, mini
    }

    return min(maxi + 1, min(n - mini, mini + 1 + n - maxi))
}
