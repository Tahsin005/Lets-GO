package main

import "fmt"

// https://leetcode.com/problems/flip-square-submatrix-vertically/description/

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	n := len(grid)
	m := len(grid[0])
	for i := x; i < x + (k / 2); i++ {
		for j := y; j < minInt(m, y + k); j++ {
			grid[i][j], grid[minInt(n, x + k) - 1 - (i - x)][j] = grid[minInt(n, x + k) - 1 - (i - x)][j], grid[i][j]
		}
	}
	return grid
}

func main () {

}