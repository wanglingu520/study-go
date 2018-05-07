package main

import "fmt"

type Man interface {
	work()
}

type China struct {
}

func (Chinese China) work() {
	fmt.Println("Hard work")
}

type Japan struct {
}

func (Japanese Japan) work() {
	fmt.Println("Smart work")
}

func main() {
	var American Man

	American = new(China)
	American.work()

	American = new(Japan)
	American.work()
}
