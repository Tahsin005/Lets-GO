package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Verteyy{3, 4}

	a = f
	a = &v

	// a = v // error: cannot use v (variable of type Verteyy) as Abser value in assignment: Verteyy does not implement Abser (Abs method has pointer receiver)

	fmt.Println(a.Abs())
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