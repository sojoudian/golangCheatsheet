package main

import "fmt"

type sc interface {
	dsc()
}

type lc interface {
	cll() int
}

type ci interface {
	sc
	lc
}

type e struct {
	f  string
	l  string
	b  int
	p  int
	tl int
	lt int
}

func (ex e) dsc() {
	fmt.Printf(" %s %s\n", ex.f, ex.l)
}

func (ex1 e) cll() int {
	return ex1.p
}

func main() {
	ex2 := e{
		"abc",
		"def",
		10,
		20,
		12,
		4,
	}

	var s sc = ex2
	s.dsc()
	var l1 lc = ex2
	fmt.Printf("%d\n", l1.cll())

	var eo ci = ex2
	eo.dsc()
	fmt.Printf(" %d\n", eo.cll())
}
