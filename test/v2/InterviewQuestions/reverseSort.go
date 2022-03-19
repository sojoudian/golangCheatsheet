package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{1, 4, 7, 9, 12, 2, 3, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
