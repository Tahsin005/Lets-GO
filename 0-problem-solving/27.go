package main

import "fmt"

// https://leetcode.com/problems/range-addition-ii/description/?envType=problem-list-v2&envId=array

func maxCount(m int, n int, ops [][]int) int {
    var min_row int = m;
    var min_col int = n;
    for i := 0; i < len(ops); i++ {
        if ops[i][0] < min_row {
            min_row = ops[i][0]
        }
        if ops[i][1] < min_col {
            min_col = ops[i][1]
        }
    }
    return min_row * min_col
}

func main () {

}