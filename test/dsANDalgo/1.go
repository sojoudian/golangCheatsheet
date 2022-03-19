package main

import "fmt"

func main() {
	hello("Albus")
	hello("Tom")
}

func hello(name string) {
	fmt.Println("Hello", name)
}
