package main

import "fmt"

func main() {
	a := []string{"a", "b"}
	fmt.Println(len(a), cap(a))
	a = append(a, "c")
	fmt.Println(len(a), cap(a))

}
