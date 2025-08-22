package main

import "fmt"

// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/description/?envType=daily-question&envId=2025-08-22

func minimumArea(grid [][]int) int {
	y1, y2 := 0, 0
	x1, x2 := 0, 0
	for i := 0; i < len(grid); i++{
		if y1 == 0 && slices.Contains(grid[i], 1){
			y1 = i + 1
		}
		if y1 != 0 && slices.Contains(grid[i], 1){
			y2 = i + 1
		}
		for j := 0; j < len(grid[i]); j++{
			if x1 == 0 && grid[i][j] == 1 {
				x1 = j + 1
			}
			if j < x1 && grid[i][j] == 1 {
				x1 = j + 1
			}
			if j + 1  > x2 && grid[i][j] == 1 {
				x2 = j + 1

			}
		}

	}
	return (x2 - x1 + 1) * (y2 - y1 + 1)
}

func main () {

}