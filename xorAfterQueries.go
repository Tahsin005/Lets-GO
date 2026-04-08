func xorAfterQueries(nums []int, queries [][]int) int {
    for _, q := range queries {
        l := q[0]
        r := q[1]
        k := q[2]
        v := q[3]
        for idx := l; idx <= r; idx += k {
            nums[idx] = (nums[idx] * v) % (1000000007)
        }
    }
    ans := 0
    for _, n := range nums {
        ans ^= n
    }
    return ans
}
