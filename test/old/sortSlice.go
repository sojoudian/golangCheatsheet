package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 122, 31, 4, 2, 3, 121, 120, 30}
	xs := []string{"Albus", "maziar", "Jinny", "Dumbledore"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(xs)
	fmt.Println(xs)

}
