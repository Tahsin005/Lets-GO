package main

import "fmt"

// https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/description/

func groupThePeople(groupSizes []int) [][]int {
    m := make(map[int][]int)
    for i, size := range groupSizes {
        m[size] = append(m[size], i)
    }

    var groups [][]int
    for size, people := range m {
        remaining := len(people)
        for remaining > 0 {
            groups = append(groups, people[remaining - size:remaining])
            remaining -= size
        }
    }

    return groups
}

func main () {

}