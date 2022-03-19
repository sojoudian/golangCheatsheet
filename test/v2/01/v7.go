package main

import "fmt"

func main() {
	fmt.Println(add(2, 3))
}

func add(a, b int) int {
	total := 0
	total = a + b
	return total
}
