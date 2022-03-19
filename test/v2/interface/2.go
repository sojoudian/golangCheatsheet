package main

import "fmt"

type I interface {
	functionOne() int
}

type St struct {
	x int
	y int
}

type Ct struct {
	a int
	b int
}

type Jt struct {
	a int
}

func (strc St) functionOne() int { //stru is a para to the method functionOne
	return strc.y + strc.x
}

func (ctrc Ct) functionOne() int { //stru is a para to the method functionOne
	return ctrc.b
}

func (ju Jt) functionOne() int {
	return ju.a
}

func functionTwo(arr []I) {
	e := 0
	for _, i := range arr {
		e = e + i.functionOne()
	}
	fmt.Printf("E = %d\n", e)
}

func main() {
	vOne := St{1, 2}
	vTwo := St{3, 4}
	cOne := Ct{5, 6}
	jOne := Jt{1}
	z := []I{vOne, vTwo, cOne, jOne}
	functionTwo(z)
}
