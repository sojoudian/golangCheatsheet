package main

import "fmt"

type A struct {
	a string
	b int
}

func methodA(vara A) {
	fmt.Printf("a: %s b: %d", vara.a, vara.b)
}

func main() {
	varB := A{
		a: "a",
		b: 10,
	}
	methodA(varB)
}
