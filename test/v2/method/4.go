package main

import "fmt"

type X struct {
	l int
	w int
}

func f(varx X) {
	fmt.Printf("l= %d, w= %d\n", varx.l, varx.w)
}

func (mvary X) methodOne() {
	fmt.Printf("mvary.l = %d, mvary.w = %d\n", mvary.l, mvary.w)
}

func main() {
	xy := X{
		14,
		17,
	}
	f(xy)
	xy.methodOne()
	v := &xy
	f(*v)
	// f(v)  // it wont work
	v.methodOne()
}
