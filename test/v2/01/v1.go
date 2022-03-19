package main

import "fmt"

var (
	m int
	n int
)

func main() {
	m = 1
	n = 2
	fmt.Println(m, n)

	a, b, c := "String", 3.14, true
	fmt.Println(a, b, c)
}
