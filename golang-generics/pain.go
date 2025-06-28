package main

import (
	"fmt"
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastItem, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	return lastItem, true
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Clear() {
	s.items = []T{}
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Peek())
	fmt.Println(intStack.Size())
	fmt.Println(intStack.IsEmpty())

	intStack.Clear()
	fmt.Println(intStack.IsEmpty())


	stringStack := Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")
	stringStack.Push("!")

	fmt.Println(stringStack.Pop())
	fmt.Println(stringStack.Pop())
	fmt.Println(stringStack.Peek())
	fmt.Println(stringStack.Size())
	fmt.Println(stringStack.IsEmpty())

	stringStack.Clear()
	fmt.Println(stringStack.IsEmpty())
}
