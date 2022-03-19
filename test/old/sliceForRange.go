package main

import (
	"fmt"
)

func main() {
	x := []int{4, 6, 7, 8, 101}
	fmt.Println(x)
	fmt.Println(len(x))
	for i, j := range x {
		fmt.Println(i, j)
	}
}
