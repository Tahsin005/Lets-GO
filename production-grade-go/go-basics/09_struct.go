package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 10
	v.Y = 20
	fmt.Println(v)

	p := &v
	fmt.Println(*p)
	p.X = 100
	fmt.Println(v)

	// ways to initialize struct
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1} // Y will be 0
	v3 := Vertex{}     // X and Y will be 0
	v4 := &Vertex{1, 2} // pointer to struct - means v4 is a pointer to Vertex
	fmt.Println(v1, v2, v3, v4)
	v4.X = 100
	fmt.Println(v4)
}