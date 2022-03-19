package main

import (
	"fmt"
)

func main() {
	//	x := make([]string ,length, cap)
	//	x := make([][]string, 50, 50)
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("reccord: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
