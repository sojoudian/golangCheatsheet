package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("len", len(a))
	fmt.Println("cap", cap(a))

	// b := []int{0, 1, 2, 3, 4, 5}
	// fmt.Println("len", len(b))
	// fmt.Println("cap", cap(b))

	b := a[1:4]

	b = b[:cap(b)]
	fmt.Println("len", len(b))
	fmt.Println(b)
}
