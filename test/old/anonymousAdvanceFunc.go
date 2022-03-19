package main

import (
	"fmt"
)

func main() {
	foo()
	func(x int) {
		fmt.Println(" my age is", x)
	}(31)
}

func foo() {
	fmt.Println(" this is not an anonymous function")
}
