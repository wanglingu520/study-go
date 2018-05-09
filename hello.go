package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64 = 3, 4

	fmt.Println(hypot(a, b))

}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
