package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("ABC", "abc"))
	fmt.Println(strings.Compare("ABC", "Abc"))
	fmt.Println(strings.Compare("abc", "Abc"))
	fmt.Println(strings.Compare("Abc", "Abc"))
}
