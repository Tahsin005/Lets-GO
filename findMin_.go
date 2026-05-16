func findMin(nums []int) int {
    l, r := 0, len(nums) - 1
	
    for l < r {
        mid := l + (r - l) / 2
        
        if nums[l] > nums[mid] { 
            r = mid 
        } else if nums[mid] > nums[r] {
            l = mid + 1
        } else if nums[mid] == nums[r] {
            r = r - 1 
        }  else {
            r = l
        }
    }
    
    return nums[l]
    
}
