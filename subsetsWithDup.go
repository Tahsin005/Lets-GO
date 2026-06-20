func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    
    subset := []int{}
    
    var dfs func(i int, subset []int)
    dfs = func(i int, subset []int) {
        if i == len(nums) {
            res = append(res, append([]int{}, subset...))
            return
        }
        dfs(i + 1, append(subset, nums[i]))
        
        for i + 1 < len(nums) && nums[i + 1] == nums[i] {
            i++
        }
        
        dfs(i + 1, subset)
    }
    
    dfs(0, subset)
    
    return res
}
