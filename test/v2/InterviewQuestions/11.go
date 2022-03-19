package main

import (
	"fmt"
)

func main() {

	const b = 10

	fmt.Println(b)

	b = 30

	fmt.Println(b)

}

// output: compilation error

// Execute go run constant.go to see the result as

// .constant.go:7:4: cannot assign to b
