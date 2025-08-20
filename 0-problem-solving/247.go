package main

import "fmt"

// https://leetcode.com/problems/count-square-submatrices-with-all-ones/description/?envType=daily-question&envId=2025-08-20

func countSquares(matrix [][]int) int {
    m, n, sum := len(matrix), len(matrix[0]), 0
    prev := make([]int, n)
    for i := 0; i < m; i++ {
        curr := make([]int, n)
        for j := 0; j < n; j++ {
            if matrix[i][j] == 1 {
                curr[j] = 1 + min(get(prev, j - 1), get(prev, j), get(curr, j - 1))
                sum += curr[j]
            }
        }
        prev = curr
    }

    return sum
}

func get(row []int, idx int) int {
    if idx < 0 {
        return 0
    }
    return row[idx]
}

func main () {

}