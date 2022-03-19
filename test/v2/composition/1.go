package main

import "fmt"

type a struct {
	f string
	l string
}

func (aElement a) methodOne() string {
	return fmt.Sprintf("%s %s", aElement.f, aElement.l)
}

type p struct {
	t string
	c string
	a
}

func (px p) methodTwo() {
	fmt.Println(px.t)
	fmt.Println(px.c)
	fmt.Println(px.methodOne())
}

func main() {
	aElement := a{"a", "b"}
	px := p{"123", "456", aElement}
	px.methodTwo()
}
