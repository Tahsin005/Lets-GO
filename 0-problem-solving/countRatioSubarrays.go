func countRatioSubarrays(nums []int, a int, b int) int {
    var result int
    for i := 0; i < len(nums); i++ {
        for j, x, y := i, 0, 0; j < len(nums); j++ {
            if nums[j] % 2 == 0 {
                x++
            } else {
                y++
            }
            if y > 0 {
                if x * b <= y * a {
                    result++
                }
            }
        }
    }
    return result
}
