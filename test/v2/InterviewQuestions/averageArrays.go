package main

import "fmt"

func main() {
	var nx [4]int
	var t, s, a int
	fmt.Println("Enter three numbers")
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		fmt.Println("Enter number of value")
		fmt.Scanln(&nx[i])
		s += nx[i]
	}
	a = s / t
	fmt.Println("The avarage of", t, " is: ", a)
}
