func diagonalSort(mat [][]int) [][]int {
    m, N, M := make(map[int][]int), len(mat), len(mat[0])
    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {
            m[i - j] = append(m[i - j], mat[i][j])
        }
    }
    for _, v := range m {
        sort.Ints(v)
    }
    for i := 0; i < N; i++ {
        for j := 0; j < M; j++ {
            mat[i][j] = m[i - j][0]
            m[i - j] = m[i - j][1:]
        }
    }
    return mat
}
