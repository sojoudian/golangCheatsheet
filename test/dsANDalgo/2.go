package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func sum(x, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	a, b := swap(10, 20)
	println(a)
	println(b)
}
