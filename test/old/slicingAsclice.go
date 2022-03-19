package main

import (
	"fmt"
)

func main() {
	x := []int{4, 6, 7, 8, 101}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[:1])
	fmt.Println(x[1:3])

	for i, j := range x {
		fmt.Println(i, j)
	}
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
