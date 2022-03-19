package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{1, 4, 7, 9, 12, 2, 3, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
	str := []string{"z", "a", "x", "b", "y"}
	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	fmt.Println(str)
}
