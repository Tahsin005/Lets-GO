func largestSubmatrix(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] > 0 {
                matrix[i][j] += matrix[i - 1][j]
            }
        }
    }
    maxA := 0
    for _, row := range matrix {
        sort.Ints(row)
        for j := n - 1; j >= 0; j-- {
            width, height := n - j, row[j]
            area := width * height
            if area > maxA {
                maxA = area
            }
        }
    }
    return maxA
}
