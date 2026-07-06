func maxDigitRange(nums []int) int {
    var result, delta int
    for i,v := range nums {
        mn, mx := v % 10, v % 10
        v /= 10
        for v > 0 {
            if v % 10 > mx {
                mx = v % 10
            }
            if v % 10 < mn {
                mn = v % 10
            }
            v /= 10
        }
        if mx - mn == delta {
            result += nums[i]
        } else if mx - mn > delta {
            result, delta = nums[i], mx - mn
        }
    }
    return result
}
