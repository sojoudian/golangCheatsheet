package main

import "fmt"

func main() {
	a := 64
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // give you the value of the certain address
	fmt.Println(*&a)
	// *a = *&a

	*b = 128
	fmt.Println(a)

}
