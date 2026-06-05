func findLengthOfLCIS(nums []int) int {
    res, current := 1, 1
    
    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i - 1] {
            current++            
        } else {
            if current > res {
                res = current
            }
            
            current = 1
        }
    }
    
    if current > res {
        res = current
    }
    
    return res
}
