package main

import "fmt"

func main() {
	// Multidimensional Array
	a := [3][2]string{{"a", "b"}, {"x", "y"}, {"q", "p"}} //3row and 2 colomns
	fmt.Println(a)
	for _, i := range a {
		for _, j := range i {
			fmt.Printf("%s\t", j)
		}
		fmt.Println("\n")
	}

	var b [2][2]string
	b[0][0] = "000"
	b[0][1] = "111"
	b[1][0] = "222"
	b[1][1] = "111"
	fmt.Println(b)
}
