package main

import "fmt"

func fn(i []int) {
	i[0] = 0
}
func main() {
	a := [3]int{89, 90, 91}
	fn(a[:])
	fmt.Println(a)
	a[1] = 1
	fmt.Println(a)
}
