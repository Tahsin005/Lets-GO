package main

import "fmt"

// https://leetcode.com/problems/matrix-cells-in-distance-order/description/?envType=problem-list-v2&envId=array

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	cells := make([][]int, 0, rows * cols)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cells = append(cells, []int{r, c})
		}
	}

	sort.Slice(cells, func(i, j int) bool {
		d1 := manhattanDistance(cells[i][0], cells[i][1], rCenter, cCenter)
		d2 := manhattanDistance(cells[j][0], cells[j][1], rCenter, cCenter)
		if d1 == d2 {
			return cells[i][0] < cells[j][0] || cells[i][0] == cells[j][0] && cells[i][1] < cells[j][1]
		}
		return d1 < d2
	})

	return cells
}

func manhattanDistance(r1, c1, r2, c2 int) int {
	return abs(r1 - r2) + abs(c1 - c2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main () {

}