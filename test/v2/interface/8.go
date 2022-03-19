package main

import "fmt"

type dd interface {
	dmethod()
}

type pp struct {
	n string
	a int
}

func (x pp) dmethod() {
	fmt.Printf("value of string %s and int: %d\n", x.n, x.a)
}

type ad struct {
	s  string
	ct string
}

func (h *ad) dmethod() {
	fmt.Printf("value of s %s and ct: %s\n", h.s, h.ct)
}

func main() {
	var dd1 dd
	pp1 := pp{
		"xyz",
		12,
	}
	dd1 = pp1
	dd1.dmethod()

	pp2 := pp{
		"abc",
		20,
	}

	//assign address of pp2  to dd1
	dd1 = &pp2
	dd1.dmethod()
	var dd2 dd
	ag := ad{
		"ttt",
		"kkk",
	}

	//dd2 = ag // will throw err because dd1 method expected pointer receiver
	dd2 = &ag
	dd2.dmethod()

}
