package main

import "fmt"

// https://leetcode.com/problems/largest-local-values-in-a-matrix/description/?envType=problem-list-v2&envId=array

func largestLocal(grid [][]int) [][]int {
	result := make([][]int, 0)
	for row := range len(grid) - 2 {
		r := make([]int, 0)
		for col := range len(grid[0]) - 2 {
			m := grid[row][col]
			for i := range 3 {
				for j := range 3 {
					m = max(m, grid[row + i][col + j])
				}
			}
			r = append(r, m)
		}
		result = append(result, r)
	}
	return result
}

func main () {

}