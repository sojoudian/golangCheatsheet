package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	fac := factorial(4)
	fmt.Println(fac)

	looop := loopFactorial(4)
	fmt.Println(looop)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
