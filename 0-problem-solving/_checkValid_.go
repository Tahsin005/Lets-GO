func checkValid(matrix [][]int) bool {
    for i := range matrix {
        freq1 := [101]int{}
        freq2 := [101]int{}
        for j := range matrix {
            freq1[matrix[i][j]]++ 
            if freq1[matrix[i][j]] > 1 {
                return false
            }
            freq2[matrix[j][i]]++ 
            if freq2[matrix[j][i]] > 1 {
                return false
            }
        }
    }
    return true
}
