package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	fac := factorial(4)
	fmt.Println(fac)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
