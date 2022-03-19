package main

import (
	"fmt"
)

func main() {
	x := []int{11, 12, 13}
	fmt.Println(x)
	x = append(x, 14, 15)
	fmt.Println(x)
	y := []int{16, 17, 18}
	fmt.Println(y)
	z := append(x, y...) //variadic
	fmt.Println(z)
	z = append(z[:2], z[4:]...)
	fmt.Println(z)

}
