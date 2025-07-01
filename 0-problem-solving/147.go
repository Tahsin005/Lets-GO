package main

import "fmt"

// https://leetcode.com/problems/number-of-employees-who-met-the-target/description/?envType=problem-list-v2&envId=array

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    count := 0
    for i := 0; i < len(hours); i++{
        if hours[i] >= target{
            count++
        }
    }
    return count
}

func main () {

}