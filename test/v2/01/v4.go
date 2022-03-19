package main

import "fmt"

const (
	x = 10
	y = 20
	z = 30
)

func main() {
	const name string = "Albus"
	fmt.Println(name)
	const a = 2
	fmt.Println(a)
	fmt.Println(x, y, z)
	// a = 4
	// fmt.Println(a)
}
