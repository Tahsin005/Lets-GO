package main

import "fmt"

// https://leetcode.com/problems/water-bottles-ii/description/?envType=daily-question&envId=2025-10-02

func maxBottlesDrunk(numBottles int, numExchange int) int {
    x := numBottles
    y := 0
    for x >= numExchange {
        x -= numExchange
        y += 1
        numExchange += 1
        x += 1
    }
    return numBottles + y
}

func main () {

}