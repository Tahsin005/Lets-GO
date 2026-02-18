package main

import "fmt"

type Struct struct {
	count int
	name string
	price float64
}

func (s *Struct) StructMethod() {
	s.count += 1
	s.name = "new keyword"
}

func NewStruct() *Struct {
	return new(Struct)
}

func main() {
	st := NewStruct()
	fmt.Printf("%T %T %T\n", st.count, st.name, st.price)
	st.StructMethod()
	fmt.Println(st.count, st.name, st.price)
}