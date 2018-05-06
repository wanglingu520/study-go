package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result := n * Factorial(n-1)
		return result
	}
	return 1
}

func main() {
	var i uint64 = 15
	//var i int =15

	fmt.Printf("%d ! = %d\n", i, Factorial(i))
	//Factorial(uint64(i))此处因为i的定义风格，与函数表达式相违背，在函数上需要进行强制转换
}
