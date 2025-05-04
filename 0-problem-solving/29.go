package main

import "fmt"

// https://leetcode.com/problems/flipping-an-image/description/?envType=problem-list-v2&envId=array

func flipAndInvertImage(image [][]int) [][]int {
    flip := func(a int) int {
        if a == 0 {
            return 1
        }
        return 0
    }

    for i := 0; i < len(image); i++ {
        l, r := 0, len(image[i]) - 1

        for l < r {
            image[i][l], image[i][r] = flip(image[i][r]), flip(image[i][l])
            l++
            r--
        }
        if l == r {
            image[i][r] = flip(image[i][r])
        }
    }

    return image
}

func main () {

}