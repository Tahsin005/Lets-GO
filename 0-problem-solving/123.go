package main

import "fmt"

// https://leetcode.com/problems/decode-xored-array/description/?envType=problem-list-v2&envId=array

func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded) + 1)
	res[0] = first
	for i := range len(encoded) {
		res[i + 1] = encoded[i] ^ res[i]
	}
	return res
}

func main() {

}