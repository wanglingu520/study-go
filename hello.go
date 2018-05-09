package main

import "fmt"

func main() {
	var a int = 7
	var b int = 8

	fmt.Println(add(a, b), sub(a, b), first(a, b), zero(a, b))
	fmt.Println(add(sub(a, b), b))

}

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}
