package main

import (
	"fmt"
	"os"
)

func main() {
	var s string = "hi"
	var sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
		// sep为空字符串，如何让其显示添加字符,给s赋值，如添加hi
	}
	fmt.Println(s)
}
