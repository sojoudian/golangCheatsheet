package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{12, 23, 13, 14, 51}
	if sort.IntsAreSorted(a) == false {
		sort.Ints(a)
	}
	fmt.Println(a)
	fmt.Println(sort.SearchInts(a, 14))

	s := []string{"zz", "yy", "xx"}
	if sort.StringsAreSorted(s) == false {
		sort.Strings(s)
	}
	fmt.Println(s)
	fmt.Println(sort.SearchStrings(s, "zz"))
}
