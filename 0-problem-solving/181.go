package main

import "fmt"

//  https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/description/

func maxWidthOfVerticalArea(points [][]int) int {
    max_diff := 0
    my_list := []int{}

    for i := 0; i < len(points); i++ {
        my_list = append(my_list, points[i][0])
    }

    sort.Ints(my_list)

    for j := 0; j < len(my_list) - 1; j++ {
        value_difference :=  my_list[j + 1] - my_list[j]
		if value_difference > max_diff {
			max_diff = value_difference
		}

    }
    return max_diff
}

func main () {

}