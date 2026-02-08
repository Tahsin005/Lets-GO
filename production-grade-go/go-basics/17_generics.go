package main

import "fmt"

func IndexInt(s []int, v int) int {
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}

func IndexString(s []string, v string) int {
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}

// generics
func Index[T comparable](s []T, v T) int {
	for i, x := range s {
		if x == v {
			return i
		}
	}
	return -1
}

func main() {
	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println(IndexInt(s, 3))

	// ss := []string{"a", "b", "c", "d", "e"}
	// fmt.Println(IndexString(ss, "f"))

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(Index(s, 3))

	ss := []string{"a", "b", "c", "d", "e"}
	fmt.Println(Index(ss, "f"))
}