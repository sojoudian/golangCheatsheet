package main

import "fmt"

func rn() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func fn(x *string, y *string) {
	defer rn()
	//defer fmt.Println("defered call in fn()")
	if x == nil {
		panic("run time err: x is nil")
	}
	if y == nil {
		panic("run time err: y is nil")
	} else {
		fmt.Printf(" %s %s", *x, *y)
	}
	fmt.Println("returning from function fn")
}

func main() {
	defer fmt.Println("defered call in main()")
	x := "abc"
	fn(&x, nil)
	fmt.Println("returning from main() function")
}
