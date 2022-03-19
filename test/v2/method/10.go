package main

import "fmt"

type A struct {
	a string
	b int
}

func (vara A) methodA() {
	fmt.Printf("a: %s \nb: %d\n", vara.a, vara.b)
}

func main() {
	varB := A{
		a: "a",
		b: 10,
	}
	varB.methodA()
}
