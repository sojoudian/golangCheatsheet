package main

import "fmt"

func main() {
	str := "Sample string"
	fmt.Printf("Before: str= %s\n", string(str))
	fmt.Printf("+++++++++++++++++++++++++++++++++++++\n")
	for _, s := range []rune(str) {
		defer fmt.Printf("%c", s)

	}

}
