func compareBitonicSums(nums []int) int {
    n := len(nums)
    
    peak := 0
    for i := 1; i < n; i++ {
        if nums[i] > nums[peak] {
            peak = i
        }
    }
    
    ascSum := 0
    for i := 0; i <= peak; i++ {
        ascSum += nums[i]
    }
    
    descSum := 0
    for i := peak; i < n; i++ {
        descSum += nums[i]
    }
    
    if ascSum > descSum {
        return 0
    } else if descSum > ascSum {
        return 1
    } 
    return -1
}
