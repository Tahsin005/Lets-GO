func numberOfSubmatrices(grid [][]byte) int {
    rows := len(grid)
    cols := len(grid[0])

    sumX := make([]int, cols)
    sumY := make([]int, cols)

    res := 0

    for i := 0; i < rows; i++ {
        rx, ry := 0, 0

        for j := 0; j < cols; j++ {
            if grid[i][j] == 'X' {
                rx++
            } else if grid[i][j] == 'Y' {
                ry++
            }

            sumX[j] += rx
            sumY[j] += ry

            if sumX[j] > 0 && sumX[j] == sumY[j] {
                res++
            }
        }
    }

    return res
}
