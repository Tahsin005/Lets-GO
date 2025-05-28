package main

import "fmt"

// https://leetcode.com/problems/matrix-diagonal-sum/description/?envType=problem-list-v2&envId=array

func diagonalSum(mat [][]int) int {
	dim := len(mat)
	maxIdx := dim - 1
	sum := 0

	for i := 0; i < dim; i++ {
		sum += mat[i][i]
		sum += mat[i][maxIdx - i]
	}

	if dim & 1 == 1 {
		sum -= mat[dim / 2][dim / 2]
	}

	return sum
}


func main () {

}