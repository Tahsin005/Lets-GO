package main

import "fmt"

// https://leetcode.com/problems/earliest-time-to-finish-one-task/

func earliestTime(tasks [][]int) int {
    if len(tasks) == 1 {
        sum := 0
        for _, t := range tasks[0] {
            sum += t
        }
        return sum
    }

    min_task := 0
    for _, t := range tasks[0] {
        min_task += t
    }

    for i := 1; i < len(tasks); i++ {
        sum := 0
        for _, t := range tasks[i] {
            sum += t
        }
        if sum < min_task {
            min_task = sum
        }
    }

    return min_task
}

func main () {

}