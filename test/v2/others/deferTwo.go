package main

import "fmt"

type p struct {
	f string
	l string
}

func (px p) methodOne() {
	fmt.Printf("%s %s\n", px.f, px.l)
}

func main() {
	px := p{
		"abc",
		"def",
	}

	defer px.methodOne()
	fmt.Printf("main functio\n")
}
