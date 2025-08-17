package main

import "fmt"

// https://leetcode.com/problems/check-if-it-is-a-straight-line/description/?envType=problem-list-v2&envId=math

func checkStraightLine(coordinates [][]int) bool {
    if len(coordinates) <= 2 {
        return true
    }


    x1, y1 := coordinates[0][0], coordinates[0][1]
    x2, y2 := coordinates[1][0], coordinates[1][1]

    dx := x2 - x1
    dy := y2 - y1


    for _, p := range coordinates[2:] {
        x, y := p[0], p[1]
        if dy * (x - x1) != dx * (y - y1) {
            return false
        }
    }
    return true
}

func main () {

}