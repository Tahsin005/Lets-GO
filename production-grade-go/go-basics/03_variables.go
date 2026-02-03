package main

import (
	"fmt"
	"math/cmplx"
)

var a, b bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var i int
	x := 4
	var y = 10
	var s string = "10"
	fmt.Println(a, b, i, x, y, s)
	fmt.Printf("%T %T %T %T %T %T\n", a, b, i, x, y, s)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
