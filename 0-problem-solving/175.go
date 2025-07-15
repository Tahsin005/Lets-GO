package main

import "fmt"

// https://leetcode.com/problems/valid-word/description/?envType=daily-question&envId=2025-07-15

func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }

    hasVowel := false
    hasConsonant := false

    for i := 0; i < len(word); i++ {
        ch := word[i]

        if ch >= '0' && ch <= '9' {
            continue
        }

        if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
            switch ch {
            case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
                hasVowel = true
            default:
                hasConsonant = true
            }

            if hasVowel && hasConsonant {
                continue
            }
        } else {
            return false
        }
    }

    return hasVowel && hasConsonant
}

func main () {

}