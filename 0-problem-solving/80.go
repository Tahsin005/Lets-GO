package main

import "fmt"

// https://leetcode.com/problems/maximum-population-year/description/?envType=problem-list-v2&envId=array

func maximumPopulation(logs [][]int) int {
    yearSet := make(map[int]int)
    ans := math.MaxInt
    for _, v := range logs {
        yearSet[v[0]] = 0
        yearSet[v[1]] = 0
    }

    for k := range yearSet {
        for _, v := range logs {
            if k >= v[0] && k <= v[1] - 1 {
                yearSet[k]++
            }
        }
    }

    maxPopulation := 0
    for year, population := range yearSet {
        if population > maxPopulation {
            maxPopulation = population
            ans = year
        } else if population == maxPopulation && year < ans {
            ans = year
        }
    }

    return ans
}

func main () {

}