package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)

	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]

	fmt.Println("###########\n", s2)
	fmt.Println(s3)
	fmt.Println(s4)

}
