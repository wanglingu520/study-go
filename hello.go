package main

import "fmt"

func main() {
	var numbers []int
	PrintSlice(numbers)

	numbers = append(numbers, 0, 1, 2, 3, 4)
	PrintSlice(numbers)

	number1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(number1, numbers)
	PrintSlice(numbers)
}

func PrintSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %d\n", len(x), cap(x), x)
}
