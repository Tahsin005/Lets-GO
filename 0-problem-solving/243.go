package main

import "fmt"

// https://leetcode.com/problems/eat-pizzas/description/

func maxWeight(pizzas []int) int64 {
    sort.Ints(pizzas)

    days := len(pizzas) / 4
    d := len(pizzas) - 1
    var ans int64

    for i := 0; i < (days + 1) / 2; i++ {
        ans += int64(pizzas[d])
        d--
    }

    d--
    for i := 0; i < days / 2; i++ {
        ans += int64(pizzas[d])
        d -= 2
    }

    return ans
}

func main () {

}