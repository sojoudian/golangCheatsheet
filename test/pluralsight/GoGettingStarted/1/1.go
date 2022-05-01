package main

import (
	"fmt"
)

func main() {
	//var digit int
	//digit = 42
	var digit int = 42
	fmt.Println(digit)

	var pi float32 = 3.14
	fmt.Println(pi)

	firstName := "Rob"
	fmt.Println(firstName)

	c := complex(3, 4)
	fmt.Println(c)

	r, i := real(c), imag(c)
	fmt.Println(r, i)
}
