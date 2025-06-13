package main

import "fmt"

// https://leetcode.com/problems/delete-greatest-value-in-each-row/description/?envType=problem-list-v2&envId=array

func deleteGreatestValue(grid [][]int) (r int) {
    for _, g := range grid {
        sort.Ints(g)
    }
    for i := len(grid[0]) - 1; i  >= 0; i--{
        x := 0
        for j := 0; j < len(grid); j++{
            if grid[j][i] > x {
                x = grid[j][i]
            }
        }

        r += x
    }
	return
}


func main () {

}