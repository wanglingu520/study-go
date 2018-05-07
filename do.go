package main

import "fmt"

func main() {
	var a int = 84
	var b int = 36

	fmt.Println(gcd(a, b))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
