package main

import "fmt"

func main() {
	i := incrementer()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	j := incrementer()
	fmt.Println(j())
	fmt.Println(j())
	fmt.Println(j())
	fmt.Println(j())
}

func incrementer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
