package main

import "fmt"

// https://leetcode.com/problems/reshape-the-matrix/description/?envType=problem-list-v2&envId=array

func matrixReshape(mat [][]int, r int, c int) [][]int {
    n, m := len(mat), len(mat[0])

    if n * m != r * c {
        return mat
    }

    arr := make([][]int, r)
    for i := 0; i < r; i++ {
		arr[i] = make([]int, c)
	}

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            l := i * m + j
            arr[l / c][l % c] = mat[i][j]
        }
    }

    return arr
}

func main () {

}