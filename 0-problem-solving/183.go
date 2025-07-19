package main

import "fmt"

// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/description/?envType=daily-question&envId=2025-07-19

import "sort"

func removeSubfolders(folder []string) []string {
    sort.Strings(folder)
    ans := []string{folder[0]}

    for i := 1; i < len(folder); i++ {
        if !startsWith(folder[i], ans[len(ans) - 1] + "/") {
            ans = append(ans, folder[i])
        }
    }
    return ans
}

// Helper function to check if a string starts with a prefix
func startsWith(s, prefix string) bool {
    if len(prefix) > len(s) {
        return false
    }
    return s[:len(prefix)] == prefix
}

func main () {

}