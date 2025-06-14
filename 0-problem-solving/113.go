package main

import "fmt"

// https://leetcode.com/problems/richest-customer-wealth/description/?envType=problem-list-v2&envId=array

func maximumWealth(accounts [][]int) int {
    for i := 0; i < len(accounts); i++ {
        for j := 1; j < len(accounts[i]); j++ {
            accounts[i][0] += accounts[i][j]
        }
        if accounts[i][0] > accounts[0][0] {
            accounts[0][0] = accounts[i][0]
        }
    }

    return accounts[0][0]
}

func main () {

}