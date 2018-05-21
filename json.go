package main

import (
	"encoding/json"
	"fmt"
)

type (
	Person struct {
		Name string `json:"nm"`
		Age  uint64 `json:"ag"`
	}
)

func main() {
	bob := Person{
		Name: "bob",
		Age:  1,
	}

	byteString, err := json.Marshal(bob)
	if err != nil {
		fmt.Printf("%v", err)

	}
	fmt.Printf("%s\n", string(byteString))

	var sam Person
	var str string = `{"nm":"bob","ag": 1}`
	err = json.Unmarshal([]byte(str), &sam)
	if err != nil {
		fmt.Printf("%v", err)

	}
	fmt.Printf("sam name: %s, age:%d", sam.Name, sam.Age)
}
