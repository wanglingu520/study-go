package main

import "fmt"

func time(a int) int {
	return (a * 12)
}

func main() {
	const manAge, womanAge = 81, 88
	fmt.Printf("%dN = %dD\n", manAge, time(manAge))
	fmt.Printf("%dN = %dD\n", womanAge, time(womanAge))
}
