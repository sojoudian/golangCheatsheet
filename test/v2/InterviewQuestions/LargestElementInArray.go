package main

import "fmt"

func main() {
	var t int
	fmt.Println("Please enter a number of values: ")
	fmt.Scanln(&t)
	fmt.Println("t is ", t)
	var nx = make([]int, t)
	//slice := make([]int, elems)

	for i := 0; i < t; i++ {
		fmt.Println("Enter the number")
		fmt.Scanln(&nx[i])
	}
	for j := 1; j < t; j++ {
		if nx[0] < nx[j] {
			fmt.Println(nx[0], "is smaller than", nx[j])
			if nx[0] < nx[j] {
				nx[0] = nx[j]
			}
		}

	}
	fmt.Println("Greatest value", nx[0])
}
