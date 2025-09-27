package main

import "fmt"

// https://leetcode.com/problems/largest-triangle-area/description/?envType=daily-question&envId=2025-09-27

type Point struct {
	X, Y float64
}

func NewPoint(x, y int) Point {
	return Point{
        X: float64(x),
        Y: float64(y),
    }
}

type Triangle struct {
	A, B, C Point
}

func (t Triangle) Area() float64 {
	return 0.5 * math.Abs(
		t.A.X * (t.B.Y - t.C.Y) +
			t.B.X * (t.C.Y - t.A.Y) +
			t.C.X * (t.A.Y - t.B.Y),
	)
}

type Geometry struct{}

func (g Geometry) MaxTriangleArea(points []Point) float64 {
	n := len(points)
	maxA := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				area := Triangle{
                    A: points[i],
                    B: points[j],
                    C: points[k],
                }.Area()
				if area > maxA {
					maxA = area
				}
			}
		}
	}

	return maxA
}

func largestTriangleArea(points [][]int) float64 {
	ps := make([]Point, len(points))
	for i, p := range points {
		ps[i] = NewPoint(p[0], p[1])
	}
	return (Geometry{}).MaxTriangleArea(ps)
}

func main () {

}