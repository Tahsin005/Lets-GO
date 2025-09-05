package main

import "fmt"

// https://leetcode.com/problems/sorting-the-sentence/description/?envType=problem-list-v2&envId=n6ww7n9h

func sortSentence(s string) string {
    resultSlice := strings.Split(s, " ")

    m := make(map[int]string)
    for _, word := range resultSlice {
        lenOfWord := len(word)
        wordIdx := int(word[lenOfWord - 1] - '0') - 1
        m[wordIdx] = string(word[:lenOfWord - 1])
    }
    var res string
    for i := 0; i < len(m); i++ {
        res += string(m[i]) + " "
    }
    return strings.TrimRight(res, " ")
}

func main () {

}