package main

import "fmt"

func main() {
	// var names []string
	// if names == nil {
	// 	names = append(names, "Albus", "Dumbledore")
	// 	fmt.Println(names)
	// }
	a := []string{"Albus", "Dumbledore"}
	b := []string{"Tom", "Marvoloe", "Riddl"}
	c := append(a, b...)
	fmt.Println(c)
	z := []int{1, 2, 3, 4, 5}
	fn(z)
	fmt.Println(z)
}

func fn(n []int) {
	for i := range n {
		n[i] = -2

	}
}
