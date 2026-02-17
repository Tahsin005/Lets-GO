func longestBalanced(nums []int) int {
    maxLen := 0

    for i := 0; i < len(nums); i++ {
        odds := make(map[int]struct{})
        evens := make(map[int]struct{})
        for j := i; j < len(nums); j++ {
            if nums[j] % 2 == 0 {
                evens[nums[j]] = struct{}{}
            } else {
                odds[nums[j]] = struct{}{}
            }
            if len(evens) == len(odds) {
                maxLen = max(maxLen, j - i + 1)      
            }
        }
    }
    return maxLen
}
