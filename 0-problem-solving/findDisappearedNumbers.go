func findDisappearedNumbers(nums []int, lower int, upper int) [][]int {
    seen := make(map[int]bool)
    for _, num := range nums {
        seen[num] = true
    }
    
    ans := [][]int{}
    idx := lower
    
    for idx <= upper {
        if !seen[idx] {
            start := idx
            for idx <= upper && !seen[idx] {
                idx++
            }
            ans = append(ans, []int{start, idx - 1})
        } else {
            idx++
        }
    }
    return ans
}
