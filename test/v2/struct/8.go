package main

import "fmt"

type n struct {
	a string
	b string
}

func main() {
	foo := n{"x", "z"}
	bar := n{"j", "v"}
	j := n{"x", "z"}

	if foo == bar {
		fmt.Println("foo and bar are equal")
	} else {
		fmt.Println("foo and bar are not equal")
	}

	if foo == j {
		fmt.Println("foo and j are equal")
	} else {
		fmt.Println("foo and j are not equal")
	}

}
