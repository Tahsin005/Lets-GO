func countGoodRotations(nums []int) int {
    n := len(nums)
    h := n / 2
    sumA, sumB := 0, 0

    for i := 0; i < h; i++ {
        sumA += nums[i]
    }

    for i := h; i < n; i++ {
        sumB += nums[i]
    }

    res  := 0
    i, x := 0, h
    for s := 0; s < n; s++ {
        if sumA > sumB {
            res++
        }
        sumA = sumA - nums[i] + nums[x]
        sumB = sumB - nums[x] + nums[i]
        i = (i + 1) % n
        x = (x + 1) % n
    }
    
    return res
}
