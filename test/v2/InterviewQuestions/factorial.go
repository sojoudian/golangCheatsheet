package main

import "fmt"

var f uint64 = 1
var iElement int = 1
var nElement int

func factorial(number int) uint64 {
	if number < 0 {
		fmt.Println("Negative number, can't  compute the factorial")
		return 0
	} else {
		for j := 1; j <= number; j++ {
			f *= uint64(j)
		}
	}
	return f
}

func main() {
	fmt.Println("Enter a positive number:")
	fmt.Scan(&nElement)
	fmt.Printf("The factorial value of the number is %d\n", factorial(nElement))
}
