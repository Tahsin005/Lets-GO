func minAbsoluteDifference(nums []int) int {
    one, two := -1, -1
    out := math.MaxInt

    for i, num := range nums {
        if num == 1 {
            one = i
            if two > -1 {
                out = min(out, one - two)
            }
        }
        if num == 2 {
            two = i
            if one > -1 {
                out = min(out, two - one)
            }
        }
    }

    if out == math.MaxInt {
        return -1
    }

    return out
}
