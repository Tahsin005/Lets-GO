package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int", v + 2)
	case string:
		fmt.Println("string", v + " world")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Verteyy{3, 4}

	a = f
	a = &v

	// a = v // error: cannot use v (variable of type Verteyy) as Abser value in assignment: Verteyy does not implement Abser (Abs method has pointer receiver)

	fmt.Println(a.Abs())

	var i I = T{"hello"}

	// only after concrete type is known, we can call the method
	i.M()

	do(21)
	do("hello")
	do(21.21)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Verteyy struct {
	X, Y float64
}

func (v *Verteyy) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}