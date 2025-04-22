package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/shortest-completing-word/?envType=problem-list-v2&envId=array

func giveShortestString(licensePlate string) map[rune]int {
    runes := []rune(strings.ToLower(licensePlate))

    mapStr := make(map[rune]int)

    for i := 0; i < len(runes); i++ {
        if (runes[i] >= 65 && runes[i] <= 90) || (runes[i] >= 97 && runes[i] <= 122) {
            mapStr[runes[i]]++
        }
    }

    return mapStr
}

func mapCheck(source, incoming map[rune]int) bool {
    ok := true

    for k, v := range source {
        if incoming[k] < v {
            ok = false
            break
        }
    }

    return ok
}

func shortestCompletingWord(licensePlate string, words []string) string {
    newLicense := giveShortestString(licensePlate)

    result := words[0]
    max := 10000
    position := 10000

    for i := 0; i < len(words); i++ {
        curr := giveShortestString(words[i])

        if mapCheck(newLicense, curr) {
            if len(words[i]) == max {
                if i < position {
                    position = i
                    max = len(words[i])
                    result = words[i]
                }
            } else if len(words[i]) < max {
                position = i
                max = len(words[i])
                result = words[i]
            }
        }
    }

    return result
}

func main () {
	
}