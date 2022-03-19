package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println(" THis is an anonymous function")
	}()

	g := func() {
		fmt.Println(" THis is a func Expression")
	}
	g()

	h := func(x int) {
		fmt.Println(" THis is an Advance func Expression", x)
	}
	h(110)

}
