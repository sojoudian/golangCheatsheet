package main

import "fmt"

type x struct {
	d map[int]int
}

func main() {
	a := x{map[int]int{0: 76}}
	b := x{map[int]int{0: 76}}

	if a == b {
		fmt.Print("a and b are equal")
	} else {
		fmt.Print("a and b are not equal")
	}

	//  this program will not work, because the struct using map is not compareable
}
