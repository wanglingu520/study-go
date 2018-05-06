package main

import "fmt"

func main() {
	var b int
	var a int

	for a = 1; a < 10; a++ {
		for b = 1; b < 10; b++ {
			fmt.Printf("%d * %d = %d\n", a, b, a*b)
		}

	}
	return
}
