func countMajoritySubarrays(nums []int, target int) int {
    var result int
    for i,_ := range nums {
        var count int
        for j := i; j < len(nums); j++ {
            if nums[j] == target {
                count++
            }
            if count * 2 > j - i + 1 {
                result++
            }
        }
    }
    return result
}
