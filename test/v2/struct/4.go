package main

import "fmt"

type Anonymous struct {
	int
	string
	string
}

func main() {
	a := Anonymous{222, "abc"}
	fmt.Println("a", a)
	fmt.Println(a.int, a.string)
}
