package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["a"] = 1
	x["b"] = 2

	fmt.Println("The map x contains", x)
	fmt.Println(x["a"])
	fmt.Println(x["b"])
}
