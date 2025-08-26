package main

import "fmt"

// https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle/description/?envType=daily-question&envId=2025-08-26

func areaOfMaxDiagonal(dimensions [][]int) int {
    maxDiagonal := 0.0
    maxArea:=0

    for i := 0; i < len(dimensions); i++ {
        length := float64(dimensions[i][0])
        width := float64(dimensions[i][1])
        diagonal := math.Sqrt(length * length + width * width)
        area := (int)(length * width)
        if (diagonal > maxDiagonal) {
            maxDiagonal = diagonal
            maxArea = area
        } else if (maxDiagonal == diagonal) {
            if area > maxArea {
                maxArea = area
            }
        }
    }

    return maxArea
}

func main () {

}