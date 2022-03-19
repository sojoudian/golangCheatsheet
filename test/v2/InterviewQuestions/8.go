package main

import "fmt"

func main() {

	a := 20

	fmt.Println("Address:", &a)

	fmt.Println("Value:", a)

}

// The result will be

// Address: 0xc000078008

// Value: 20
