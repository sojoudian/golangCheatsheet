package main

import "fmt"

func fOne(i int) {
	fmt.Println("i= ", i)
}

func main() {
	n := 10
	defer fOne(n)
	n = 20
	fmt.Println("n= ", n)
}
