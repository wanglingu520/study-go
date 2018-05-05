package main

import "fmt"

func main() {

	var b int = 15
	var a int

	numbers := [8]int{1, 2, 3, 7}

	for a = 0; a < 20; a++ {
		fmt.Printf("%d\n", a)
	}

	for a < b {
		fmt.Printf("%d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("%d x = %d\n", i, x)
	}
}
