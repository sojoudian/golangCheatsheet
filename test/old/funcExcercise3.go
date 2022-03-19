package main

import "fmt"

func main() {
	foo()
	fmt.Println("After foo")
}

func foo() {
	defer func() {
		fmt.Println("Anonymous func to DEFER run")
	}()
	fmt.Println("Foo runned")
}
