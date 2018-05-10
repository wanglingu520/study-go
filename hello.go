package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	var a = Point{7, 6}
	var b = Point{5, 4}
	fmt.Printf("p to q distance = %f\n", Distance(a, b))
	fmt.Println(a.Distance(b))
}
