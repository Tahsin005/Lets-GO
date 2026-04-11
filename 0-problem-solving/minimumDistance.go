func minimumDistance(nums []int) int {
    var result int = 1000000
    m := make(map[int][]int)
    for i, v := range nums {
        m[v] = append(m[v], i)
    }

    for _, w := range m {
        if len(w) >= 3 {
            for i := 2; i < len(w); i++ {
                if w[i] - w[i - 2] < result {
                    result = w[i] - w[i - 2]
                }
            }
        }
    }

    if result == 1000000 {
        return -1
    }
    
    return result * 2
}
