package main

import (
	"fmt"
)

// type person struct{
// 	first string
// 	last string
// 	age int
// }

func main() {
	p := struct {
		first string
		last  string
		age   int
	}{
		first: "ken",
		last:  "Thompson",
		age:   78,
	}
	fmt.Println(p)
}
