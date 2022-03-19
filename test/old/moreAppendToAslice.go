package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	for i := 53; i <= 55; i++ {
		x = append(x, i)
	}
	fmt.Println(x)
	x = append(x, 56, 57, 58)
	fmt.Println(x)
	y := []int{59, 60, 61, 62, 63}
	z := append(x, y...)
	fmt.Println(z)
}
