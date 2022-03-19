package main

import "fmt"

type X struct {
	a string
	b int
}

func (x X) methodX() {
	fmt.Printf("%s %d\n", x.a, x.b)
}

type Y struct {
	xy string
	ab string
	X  // this X is Anonymous structure
}

func main() {
	jk := Y{
		"abc",
		"123",
		X{
			"jkl",
			567,
		},
	}
	jk.methodX()
}
