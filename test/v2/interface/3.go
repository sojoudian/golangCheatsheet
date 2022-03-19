package main

import "fmt"

type I interface {
	funnctionOne()
}

type fLoat float64

func (flVal fLoat) funnctionOne() { //method
	fmt.Println("\nvalue is ", flVal)
}

func functionTwo(i I) {
	fmt.Printf("type %T value %v", i, i)
}

func main() {
	var a I
	k := fLoat(10.2)
	a = k
	functionTwo(a)
	fmt.Println("\n>>")
	a.funnctionOne()
}
