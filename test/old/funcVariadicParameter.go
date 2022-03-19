package main

import (
	"fmt"
)

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total number is: ", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Printf("for item  in index position", i, "we are now adding", v, "value for the total number is now", sum)
	}
	fmt.Printf("The total is: ", sum)
	return sum
}
