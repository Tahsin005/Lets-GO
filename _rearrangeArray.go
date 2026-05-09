func rearrangeArray(nums []int) []int {
    ans := make([]int, len(nums))
    var pos, neg []int

    for _, num := range nums {
        if num > 0 {
            pos = append(pos, num)
        } else {
            neg = append(neg, num)
        }
    }

    for i := 0; i < len(pos); i++ {
        ans[i * 2] = pos[i]
        ans[i * 2 + 1] = neg[i]
    }

    return ans
}
