package main

import "fmt"

func test() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	PrintSlice(numbers)
	fmt.Println("numbers==", numbers)
	number1 := numbers[1:4]
	PrintSlice(number1)
	fmt.Println("number[:3] ==", numbers[:3])
	number2 := numbers[4:]
	PrintSlice(number2)
	number3 := make([]int, 0, 5)
	PrintSlice(number3)
	number4 := numbers[:2]
	PrintSlice(number4)
}

func PrintSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
