package main

import "fmt"

// https://leetcode.com/problems/water-bottles/description/?envType=daily-question&envId=2025-10-01

func numWaterBottles(numBottles int, numExchange int) int {
    cntBottle := numBottles
    khaliBottle := numBottles

    for khaliBottle > 0 {
        if khaliBottle >= numExchange {
            khaliBottle -= numExchange
            cntBottle++
            khaliBottle++
        } else {
            break
        }
    }

    return cntBottle
}

func main () {

}