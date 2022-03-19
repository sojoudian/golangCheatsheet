package main

import (
	"fmt"
)

func main() {
	foo()
	func() {
		fmt.Println(" this is an anonymous function")
	}()
}

func foo() {
	fmt.Println(" this is not an anonymous function")
}
