package main

import (
	"fmt"
	"s/subModule"
)

func main() {
	var a subModule.S
	a.M = "this is a string"

	fmt.Println("a is:", a)
}
