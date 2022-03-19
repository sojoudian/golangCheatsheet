package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 10, 2, 20, 30, 3}
	xs := []string{"X", "y", "Y", "Z", "z"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}
