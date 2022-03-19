package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	//aa := []int{1, 2, 3, 4, 5}
	fmt.Println("original array", a)

	s := a[1:3]
	fmt.Println("sliced array", s)

	for i := range s {
		s[i]++
	}
	fmt.Println("original array after modification", a)
	fmt.Println("sliced array after modification", s)

	r := a[1:4]
	for i := range r {
		r[i]++
	}

	fmt.Println("original array after second modification", a)
	fmt.Println("sliced array after modification", r)
	fmt.Println(len(r))
}
