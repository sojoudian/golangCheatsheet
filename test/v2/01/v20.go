package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	var b = a[1:4]
	fmt.Println(b)
}
