package main

import (
	"fmt"
)

func main() {
	u1 := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(u1)
	u2 := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(u2)

	combine := [][]string{u1, u2}
	fmt.Println(combine)

}
