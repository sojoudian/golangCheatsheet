package main

import "fmt"

func funcOne() {
	fmt.Println("print from funcOne")
}

func funcTwo(n []int) {
	defer funcOne()
	fmt.Println("executing funcTwo")
	m := n[0]
	for _, val := range n {
		if val > m {
			m = val
		}
	}
	fmt.Println("Has got", m, "as the biggest number")
}

func main() {
	n := []int{6, 5, 4, 3, 2, 1}
	funcTwo(n)
}
