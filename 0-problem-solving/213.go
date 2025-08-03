package main

import "fmt"

// https://leetcode.com/problems/earliest-finish-time-for-land-and-water-rides-i/description/

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    n := len(landStartTime); n2 := len(waterDuration)
    arr1 := make([]Pair, n); arr2 := make([]Pair, n2)
    for i := 0; i < n; i++ {
        arr1[i] = Pair{v: landStartTime[i] + landDuration[i], k: i}
    }
    for i := 0; i < n2; i++ {
        arr2[i] = Pair{v: waterStartTime[i] + waterDuration[i],k: i}
    }
    sort.Slice(arr1, func(i, j int)bool {
        return arr1[i].v < arr1[j].v
    })
    sort.Slice(arr2, func(i, j int)bool {
        return arr2[i].v < arr2[j].v
    })

    res1 := 10000
    for i := 0; i < n2; i++ {
        min := Max(arr1[0].v, waterStartTime[i])
        min += waterDuration[i]
        res1 = Min(res1, min)
    }
    res2 := 10000
    for i := 0; i < n; i++ {
        min := Max(arr2[0].v, landStartTime[i])
        min += landDuration[i]
        res2 = Min(res2, min)
    }
    return Min(res1, res2)
}

func Max(a, b int)int {
    if a < b {
        return b
    }
    return a
}

func Min(a, b int)int {
    if a > b {
        return b
    }
    return a
}
type Pair struct {
    v int
    k int
}

func main () {

}