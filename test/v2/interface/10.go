package main

import "fmt"

type zI interface {
	dMethod()
}

func main() {
	var zIOne zI
	if zIOne == nil {
		fmt.Printf("%T %v", zIOne, zIOne)
	}

	// zIOne.dMethod() // will throw an error
}
