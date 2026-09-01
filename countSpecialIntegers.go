func countSpecialIntegers(nums []int) int {
    m := make(map[int]int)
    m[nums[0]] = 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i - 1] {
            continue
        }
        m[nums[i]]++
    }
    
    var result int
    for _,v := range m {
        if v == 1 {
            result++
        }
    }
    return result
}
