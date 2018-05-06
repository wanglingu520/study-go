package main

import "fmt"

type Phone interface {
	call()
}

type Nokiaphone struct {
}

func (nokiaphone Nokiaphone) call() {
	fmt.Println("Nokia call")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("Iphone call")
}
func main() {
	var phone Phone

	phone = new(Nokiaphone)
	phone.call()

	phone = new(Iphone)
	phone.call()
}
