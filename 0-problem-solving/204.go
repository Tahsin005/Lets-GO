package main

import "fmt"

// https://leetcode.com/problems/rings-and-rods/description/?envType=problem-list-v2&envId=hash-table

func countPoints(rings string) int {
    rods := make([][]int, 10)

    for i := 0; i < 10; i++ {
        rods[i] = make([]int, 3)
    }

    for i := 0; i < len(rings); i += 2 {
        c := rings[i]
        rod := rings[i + 1]

        if c == 'B' {
            rods[rod - '0'][0] = 1
        } else if c == 'G' {
            rods[rod - '0'][1] = 1
        } else if c == 'R' {
            rods[rod - '0'][2] = 1
        }
    }

    c := 0
    for i := 0; i < 10; i++ {
        if rods[i][0] == 1 && rods[i][1] == 1 && rods[i][2] == 1 {
            c++
        }
    }
    return c
}

func main () {

}