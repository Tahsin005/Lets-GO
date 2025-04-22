package main

import "fmt"

func twoSum(nums []int, target int) []int {
    counts := make(map[int]int)
    for i, num := range nums {
        required := target - num
        index, found := counts[required]
        if found {
            return []int{index, i}
        }
        counts[num] = i
    }
    return nil
}

func main () {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("=====================================")
	fmt.Println(twoSum([]int{3, 2, 4}, 10))
	fmt.Println("=====================================")
	fmt.Println(twoSum([]int{3, 3}, 6))
}