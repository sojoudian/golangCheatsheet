package main

import "fmt"

type dm interface {
	dMethod()
}

type multipleStruct struct {
	s string
	n int
}

func (mpElement multipleStruct) dMethod() {
	fmt.Printf("String is %s, and int os %d\n", mpElement.s, mpElement.n)
}

func fOne(i interface{}) {
	switch v := i.(type) {
	case dm:
		v.dMethod()
	default:
		fmt.Printf("unkown type of value\n")
	}
}

func main() {
	fOne("Tom")
	k := multipleStruct{
		"Albus",
		106,
	}
	fOne(k)
}
