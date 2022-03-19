// What is the output of the below pgm:

package main

import "fmt"

func fact(n int) int {

	if n == 0 {

		return 1

	}

	return n * fact(n-1)

}

func main() {

	fmt.Println(fact(7))

}

//This fact function calls itself until it reaches the base case of fact(0).

// $ go run recursion.go

// 5040
