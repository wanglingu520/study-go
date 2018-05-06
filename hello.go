package main

func main() {
	var numbers []int
	PrintSlice(numbers)

	numbers = append(numbers, 0, 1, 2, 3, 4)
	PrintSlice(numbers)

	number1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(number1, numbers)
	PrintSlice(numbers)
}
